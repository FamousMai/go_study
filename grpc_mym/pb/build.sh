#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello_grpc.proto
#protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./person/person.proto

## 以gateway的方式去生成
protoc --go_out=. --go_opt paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative ./person/person.proto

  ## todo 执行后报以下错误
  ## protoc-gen-grpc-gateway: program not ecfound or is not executable
     #Please specify a program using absolute path or make sure the program is available in your PATH system variable
     #--grpc-gateway_out: protoc-gen-grpc-gateway: Plugin failed with status code 1.
  ## todo 执行go install之后 就没问题了
  ## go install \
       #      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
       #      github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
       #      google.golang.org/protobuf/cmd/protoc-gen-go \
       #      google.golang.org/grpc/cmd/protoc-gen-go-grpc
