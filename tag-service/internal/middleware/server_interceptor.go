package middleware

import (
	"context"
	"log"
	"runtime/debug"
	"time"

	"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	"google.golang.org/grpc"
)

func AccessLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestLog := "access request log: method :%s, begin_time:%d, request:%v"
	begin_time := time.Now().Local().Unix()
	log.Printf(requestLog, info.FullMethod, begin_time, req)

	resp, err := handler(ctx, req)

	responseLog := "access response log : method :%s, begin_time: %d, end_time: %d, response: %v"
	end_time := time.Now().Local().Unix()

	log.Printf(responseLog, info.FullMethod, begin_time, end_time, resp)
	return resp, err
}

func ErrorLog(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)

	if err != nil {
		errLog := "error log: method :%s, code :%v, message :%v, details %v"
		s := errcode.FromError(err)
		log.Printf(errLog, info.FullMethod, s.Code(), s.Err().Error(), s.Details())
	}
	return resp, err
}

func Recovery(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	defer func() {
		if e := recover(); e != nil {
			recoveryLog := "recovery log: method: %s, message: %v, stack: %s"
			log.Printf(recoveryLog, info.FullMethod, e, string(debug.Stack()[:]))
		}
	}()

	return handler(ctx, req)
}
