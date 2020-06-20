#install and start minikube
minikube:
		brew install kubectl
		brew install minikube

start:
		minikube start
		#minikube addons enable ingress
#stop docker and minikube. remove exsisting container
clean:
		docker stop go-app-container
		docker rm go-app-container
#delete all the pods and deployments in the kubernetes cluster
		kubectl delete --all pods --namespace=default
		kubectl delete --all svc --namespace=default
		kubectl delete --all deployments --namespace=default
		#gcloud container clusters delete test1
stop:
		minikube stop
#build the docker image using the docker file in the current directory
#create a docker container locally and expose it on the port 8080
build:
		docker build -t go-app-img .
		docker run -d -p 8080 --name go-app-container go-app-img
#this is required if we are deploying on gks
gcloud:
		brew cask install google-cloud-sdk
		gcloud init
#push to docker public registry
push:
		docker login
		docker build -t golangrest/go-app .
		docker push golangrest/go-app:latest

#provision resources if we are deploying on a gks clusters
provision:
		gcloud container clusters create test1 --zone us-east1-b
		gcloud container clusters get-credentials test1 --zone us-east1-b --project profilescreeningagent

deploy:
		make build push
		kubectl create -f deployment.yaml
		#kubectl get deployments
		kubectl expose deployment my-go-app --type=NodePort --name=go-app-svc --port=3000
		kubectl get pods
		#kubectl get service my-go-app


portforward:
		kubectl port-forward my-go-app-7d7cf64f76-d8rjs 7000:7000

test:

		curl localhost:7000/people/1
