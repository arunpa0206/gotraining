docker-install-windows:
		choco install docker-desktop

docker-install-mac:
		brew cask install docker

minikube-install-windows:
		choco install minikube
		
minikube-install-mac:
		brew install kubectl
		brew install minikube

aws-install-windows:
		choco install awscli
		choco install eksctl
		aws configure

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

docker-clean:
		docker stop $(container-name)
		docker rm $(container-name)

docker-clean-test:	
		make container-name=go-app-container docker-clean
		make image-name=go-app docker-build
		make container-name=go-app-container image-name=go-app docker-run
		make url=localhost:8030/people/1 test-docker

docker-test:
		make image-name=go-app docker-build
		make container-name=go-app-container image-name=go-app docker-run
		make url=localhost:8030/people/1 test-docker

minikube-start:
		minikube start

aws-start:
		eksctl create cluster --name $(cluster-name) --region $(region)
		aws eks --region $(region) update-kubeconfig --name $(cluster-name)

deploy:
		make minikube-aws-clean
		kubectl create -f deployment.yaml
		kubectl expose deployment $(app-name) --type=NodePort --name=$(svc) --port=3000
		kubectl get pods

minikube-prepare:
		make minikube start
		make repo=golangrest image-name=go-app tag=latest docker-push
		make app-name=my-go-app svc=go-app-svc deploy

aws-prepare:
		make aws-start
		make repo=golangrest image-name=go-app tag=latest docker-push
		make app-name=my-go-app svc=go-app-svc deploy

portforward:
		kubectl port-forward $(pod) 7000:7000

test:
		curl $(url)
		
minikube-aws-clean:
		kubectl delete --all pods --namespace=$(namespace)
		kubectl delete --all svc --namespace=$(namespace)
		kubectl delete --all deployments --namespace=$(namespace)

minikube-testandstop:
		make url=localhost:7000/people/1 minikube-test
		minikube-stop

aws-testandstop:
		make url=localhost:7000/people/1 aws-test
		eksctl delete cluster --name $(cluster-name) --region $(region)





























