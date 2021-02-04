docker-install-windows:
		choco install docker-desktop

docker-install-mac:
		brew cask install docker
		

minikube-install-windows:
		choco install minikube
		

minikube-install-mac:
		brew install kubectl
		brew install minikube

install-windows:
		make docker-install-windows
		make minikube-install-windows

install-mac:
		
		make docker-install-mac
		make minikube-install-mac


docker-clean:
		docker stop $(container-name)
		docker rm $(container-name)


docker-build:
		docker build -t $(image-name) .

docker-run:
		docker run -d -p 8030:7000 --name $(container-name)  $(image-name)


test-docker:
		curl $(url)


docker-push:
		docker login
		docker build -t $(repo)/$(image-name) .
		docker push $(repo)/$(image-name):$(tag)


minikube-start:
		minikube start
		
minikube-clean:
		kubectl delete --all pods --namespace=$(namespace)
		kubectl delete --all svc --namespace=$(namespace)
		kubectl delete --all deployments --namespace=$(namespace)

minikube-deploy:
		kubectl create -f deployment.yaml
		kubectl expose deployment $(app-name) --type=NodePort --name=$(svc) --port=3000
		kubectl get pods

minikube-portforward:
		kubectl port-forward $(pod) 7000:7000

minikube-test:
		curl $(url)

minikube-stop:
	minikube stop


docker-clean-test:
	
		make container-name=go-app-container docker-clean
		make image-name=go-app docker-build
		make container-name=go-app-container image-name=go-app docker-run
		make url=localhost:8030/people/1 test-docker

docker-test:
		make image-name=go-app docker-build
		make container-name=go-app-container image-name=go-app docker-run
		make url=localhost:8030/people/1 test-docker

prepare:
		
		make minikube-start
		make namespace=default minikube-clean
		make docker-clean-test
		make repo=golangrest image-name=go-app tag=latest docker-push
		make app-name=my-go-app svc=go-app-svc minikube-deploy

minikube-testandstop:
		make url=localhost:7000/people/1 minikube-test
		make minikube-stop
