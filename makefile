get:
		go get github.com/gorilla/mux
		go get google.golang.org/protobuf/cmd/protoc-gen-go
		go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go get -u google.golang.org/grpc
		go get -u github.com/golang/protobuf/protoc-gen-go
		go get github.com/gocql/gocql
		go get github.com/go-sql-driver/mysql
		go get github.com/newrelic/go-agent
		go get golang.org/x/net/context
		go get github.com/streadway/amqp
		go mod init
		go get github.com/go-redis/redis/v7
		go get gopkg.in/check.v1






	

install-windows:
		
		cd  14.b.kubernetes && make install-windows
		cd 15.microservices && make install-windows
		cd 17.db && make install-windows
		cd 24.messaging && make install-windows
		make get
 
 install-mac:
		
		cd  14.b.kubernetes && make install-mac
		cd 15.microservices && make install-mac
		cd 17.db && make install-mac
		cd 24.messaging && make install-mac
		make get

test:
		protoc --vesrion
		make --vesrion
		rabbitmqctl --version
		docker --version
		minikube --version
		mysql --version
