Install  minikube:

https://kubernetes.io/docs/tasks/tools/install-minikube/

Install make:
brew install make

inorder to run the service on kubernetes:
 make start
 make deploy
 copy the pod name and change it in portforward target
 make portforward
 make test
 make clean stop
