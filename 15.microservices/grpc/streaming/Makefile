PROTO_PATH := api/proto

# Output directories.
GRPC_OUT := api/generated

protoc:
	# Generate proto stubs.
	protoc \
	-I $(PROTO_PATH) \
	--go_out=plugins=grpc:$(GRPC_OUT) \
	$(PROTO_PATH)/*.proto

run:
	cd server && go run server.go


unary-client:
	cd client && go run unary-client.go
	cd ..

client-streaming:
	cd client && go run clientstreaming.go
	
server-streaming:
	cd client && go run serverstreaming-client.go

bidir-streaming:
	cd client && go run bisir-streaming-client.go



