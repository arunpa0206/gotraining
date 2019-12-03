 Install grpc --
 go get -u google.golang.org/grpc

 Install protoc --
 go get -u github.com/golang/protobuf/protoc-gen-go

 To compile proto file 

 protoc --go_out=plugins=grpc:. hello.proto 
