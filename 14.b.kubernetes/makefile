docker-clean:
		docker stop go-app-container
		docker rm go-app-container

#build the docker image using the docker file in the current directory
#create a docker container locally and expose it on the port 8080
docker-build:
		docker build -t go-app-img .
docker-run:
		docker run -d -p 8030:7000 --name go-app-container  go-app-img
#this is required if we are deploying on gks

test-docker:
		curl localhost:8030/people/1

#gcloud:
		#brew cask install google-cloud-sdk
		#gcloud init

#push to docker public registry
docker-push:
		docker login
		docker build -t golangrest/go-app .
		docker push golangrest/go-app:latest
#install and start minikube
minikube-install:
		brew install kubectl
		brew install minikube

minikube-start:
		minikube start
		#minikube addons enable ingress
#stop docker and minikube. remove exsisting container

minikube-clean:
		#delete all the pods and deployments in the kubernetes cluster
		kubectl delete --all pods --namespace=default
		kubectl delete --all svc --namespace=default
		kubectl delete --all deployments --namespace=default
		#gcloud container clusters delete test1

#provision resources if we are deploying on a gks clusters
#provision:
		#gcloud container clusters create test1 --zone us-east1-b
		#gcloud container clusters get-credentials test1 --zone us-east1-b --project profilescreeningagent

minikube-deploy:
		#make build push
		kubectl create -f deployment.yaml
		#kubectl get deployments
		kubectl expose deployment my-go-app --type=NodePort --name=go-app-svc --port=3000
		kubectl get pods
		#kubectl get service my-go-app


minikube-portforward:
		kubectl port-forward my-go-app-7d7cf64f76-qbnp8 7000:7000

minikube-test:

		curl localhost:7000/people/1

minikube-stop:
		minikube stop


docker-test:
	
		make docker-clean
		make docker-build
		make docker-run
		make test-docker

prepare:
		#make minikube-install
		make minikube-start
		make minikube-clean
		make docker-test
		make docker-push
		make minikube-deploy

minikube-testandstop:
		make minikube-test
		make minikube-stop
