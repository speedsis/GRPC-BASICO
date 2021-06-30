gen:
	protoc --go_out=framework/grpc/pb --go_opt=paths=source_relative --go-grpc_out=framework/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=framework/grpc/protofiles framework/grpc/protofiles/*.proto

clean:
	rm pb/*.go

server:
	go run framework/cmd/server/main.go
	#go run framework/cmd/server/main.go -port 8080

client:
	go run framework/cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...

evans:
	evans -r repl
	#evans -r repl -p 8080

.PHONY: gen clean server client test evans