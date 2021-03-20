rabbitmq:
		go run rabbitmq/send.go	
		go run rabbitmq/receive.go
		
redis-send:
		go run redis/client.go redis/broker.go 

redis-recieve:
		go run redis/worker.go redis/broker.go

install-windows:
		choco install rabbitmq
		rabbitmqctl status
		

install-mac:
		brew update
		brew install rabbitmq
		export PATH=$PATH:/usr/local/sbin
		brew services start rabbitmq
		rabbitmqctl status
		brew install redis
		brew services start redis
		redis-cli ping



           		
		