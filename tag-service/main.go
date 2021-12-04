package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/proxy/grpcproxy"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/go-programming-tour-book/tag-service/internal/middleware"
	"github.com/go-programming-tour-book/tag-service/pkg/swagger"
	pb "github.com/go-programming-tour-book/tag-service/proto"
	"github.com/go-programming-tour-book/tag-service/server"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// var grpcPort string
// var httpPort string

// ------------------- Another port listen http request ---------------//

// func init() {
// 	flag.StringVar(&grpcPort, "grpc_port", "8008", "grpc port")
// 	flag.StringVar(&httpPort, "http_port", "9001", "http port")
// 	flag.Parse()
// }

// func RunHttpServer(port string) error {
// 	serverMux := http.NewServeMux()
// 	serverMux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
// 		_, _ = rw.Write([]byte(`pong`))
// 	})

// 	return http.ListenAndServe(":"+port, serverMux)
// }

// func RunGrpcServer(port string) error {
// 	s := grpc.NewServer()
// 	pb.RegisterTagServiceServer(s, server.NewTagServer())
// 	reflection.Register(s) // grpcurl test

// 	lis, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		return err
// 	}
// 	return s.Serve(lis)
// }

// func main() {
// 	errs := make(chan error)

// 	go func() {
// 		err := RunHttpServer(httpPort)
// 		if err != nil {
// 			errs <- err
// 		}
// 	}()

// 	go func() {
// 		err := RunGrpcServer(grpcPort)
// 		if err != nil {
// 			errs <- err
// 		}
// 	}()

// 	select {
// 	case err := <-errs:
// 		log.Fatalf("Run Server err:%v", err)
// 	}
// }

// --------------------Using Cmux to listen http request and grpc request in same port --------//
// var port string

// func init() {
// 	flag.StringVar(&port, "port", "8008", "init port")
// 	flag.Parse()
// }

// func RunTCPServer(port string) (net.Listener, error) {
// 	return net.Listen("tcp", ":"+port)
// }

// func RunGrpcServer() *grpc.Server {
// 	s := grpc.NewServer()
// 	pb.RegisterTagServiceServer(s, server.NewTagServer())
// 	reflection.Register(s)
// 	return s
// }

// func RunHttpServer(port string) *http.Server {
// 	serveMux := http.NewServeMux()
// 	serveMux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
// 		_, _ = rw.Write([]byte("sac"))
// 	})

// 	return &http.Server{
// 		Addr:    ":" + port,
// 		Handler: serveMux,
// 	}
// }

// func main() {
// 	l, err := RunTCPServer(port)
// 	if err != nil {
// 		log.Fatalf("Run TCP Server Error:%v", err)
// 	}

// 	m := cmux.New(l)
// 	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
// 	httpL := m.Match(cmux.HTTP1Fast())

// 	grpcS := RunGrpcServer()
// 	httpS := RunHttpServer(port)
// 	go grpcS.Serve(grpcL)
// 	go httpS.Serve(httpL)

// 	err = m.Serve()
// 	if err != nil {
// 		log.Fatalf("Run Server err : %v", err)
// 	}

// }

// ------------------ Http And RPC  + interpretor ------------------//
var port string

const SERVICE_NAME = "tag-service"

func init() {
	flag.StringVar(&port, "port", "8004", "init port")
	flag.Parse()
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	// h2c.NewHandler()
	return h2c.NewHandler(
		http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
				grpcServer.ServeHTTP(rw, r)
			} else {
				otherHandler.ServeHTTP(rw, r)
			}
		}), &http2.Server{})

}

func RunServer(port string) error {
	// return nil
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer()
	httpMux.Handle("/", gatewayMux)
	runtime.HTTPError = grpcGatewayError

	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: time.Second * 60,
	})

	if err != nil {
		return err
	}

	defer etcdClient.Close()
	target := fmt.Sprintf("/etcdv3://go-programming-tour/grpc/%s", SERVICE_NAME)
	grpcproxy.Register(etcdClient, target, ":"+port, 60)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		_, _ = rw.Write([]byte("Helloword"))
	})

	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})

	serverMux.Handle(prefix, http.StripPrefix(prefix, fileServer))

	serverMux.HandleFunc("/swagger/", func(rw http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(rw, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)
		http.ServeFile(rw, r, p)
	})
	return serverMux
}

func runGrpcServer() *grpc.Server {

	opts := []grpc.ServerOption{
		// grpc.UnaryInterceptor(HelloInterceptor),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)
	return gwmux
}

func main() {
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// grpc error trans to http error
func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()

	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}

func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("hello")
	resp, err := handler(ctx, req)
	log.Println("GoodBye")
	return resp, err
}

func WorldInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Travel Around the World")
	resp, err := handler(ctx, req)
	log.Println("Failed because of poverty")
	return resp, err
}
