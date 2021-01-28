run:
		go run json/test_json.go
		make proto-run
	

build:
		protoc --go_out=proto/. proto/basic.proto

proto-run:
		go run proto/main.go proto/basic.pb.go

		
