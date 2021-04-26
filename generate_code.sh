#L16 Code Generation Test
# This doesn't work anymore. (protoc version: libprotoc 3.15.8)
#refer https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code for the up-to-date command
protc greet/greetpb/greet.proto --go_out=plugins=grpc:.

#This command works just fine
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet/greetpb/greet.proto