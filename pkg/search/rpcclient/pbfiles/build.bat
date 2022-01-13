protoc --proto_path=. --plugin=protoc-gen-go=D:\develop\GolandProjects\gopath\bin\protoc-gen-go.exe --go_out=./ apis.proto

protoc --proto_path=. --plugin=protoc-gen-go=D:\develop\GolandProjects\gopath\bin\protoc-gen-go.exe --go-grpc_out=./  apis.proto
