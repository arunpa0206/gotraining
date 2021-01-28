

# gotraining

1. Installation Guide 
https://golang.org/doc/install 

2. Set Up  Workspace

https://www.callicoder.com/golang-installation-setup-gopath-workspace/ 

3.Install mysql

https://gist.github.com/nrollr/3f57fc15ded7dddddcc4e82fe137b58e

4. Signup for newrelic free account on the new relic website and have a login

5.Install minikube

https://kubernetes.io/docs/tasks/tools/install-minikube/

6.Install docker desktop:

https://docs.docker.com/docker-for-windows/install/

7.Set environment

		export GOPATH=~/go
  		export GOBIN=$GOPATH/bin
		export PATH=$PATH:$GOPATH:$GOBIN

8.Install redis -mac
https://medium.com/@petehouston/install-and-config-redis-on-mac-os-x-via-homebrew-eb8df9a4f298  

Installing redis on windows:
https://redislabs.com/ebook/appendix-a/a-3-installing-on-windows/a-3-2-installing-redis-on-window/ 

stalling RabbitMQ: MAC
                brew update
                brew install rabbitmq
                export PATH=$PATH:/usr/local/sbin
                brew services start rabbitmq
                rabbitmqctl status

Installing rabbitMQ: windows
            Download Erlang : https://www.erlang.org/downloads 
            Download Rabbitmq: https://www.rabbitmq.com/install-windows.html#downloads 

Installing kafka on windows:
https://dzone.com/articles/running-apache-kafka-on-windows-os 
        
Installing Kafka on mac:
        brew cask install java
        brew install kafka
        zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
        kafka-server-start /usr/local/etc/kafka/server.properties
9.Install make for mac:
brew install make

install make for windows:
https://stackoverflow.com/questions/32127524/how-to-install-and-use-make-in-windows - windows

Install Protobuf:
 brew install protobuf -  MAC
 https://github.com/protocolbuffers/protobuf/releases - windows
 
 