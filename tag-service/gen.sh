protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:. --grpc-gateway_out=:pb 



protoc -I/usr/local/include -I. \
       -I$GOPATH/src \
       -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis\
       --grpc-gateway_out=logtostderr=true:. \
       ./proto/*.proto

protoc -I .\
       -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.5/third_party/googleapis\
       --grpc-gateway_out=logtostderr=true:. \
       --swagger_out=logtostderr=true:.\
       ./proto/*.proto


# failed 
protoc -I. \
  -I ./proto \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
  ./proto/tag.proto ./proto/comman.proto