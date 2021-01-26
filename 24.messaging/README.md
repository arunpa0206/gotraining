Installing RabbitMQ: MAC
                brew update
                brew install rabbitmq
                export PATH=$PATH:/usr/local/sbin
                brew services start rabbitmq
                rabbitmqctl status

Installing rabbitMQ: windows
            Download Erlang : https://www.erlang.org/downloads 
            Download Rabbitmq: https://www.rabbitmq.com/install-windows.html#downloads 

Installing go driver for RabbitMQ:
                go get github.com/streadway/amqp


Installing redis:
            brew install redis
            brew services start redis
            redis-cli ping

Installing go driver for redis:
            go mod init github.com/my/repo
            go get github.com/go-redis/redis/v8