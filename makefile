generate-pb:
	protoc --go_out=. --go-grpc_out=. stock.proto