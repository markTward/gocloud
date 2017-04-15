#TODO
travis deploy using custom scripts: https://docs.travis-ci.com/user/deployment/script
Split apart grpc and restapi into different repos/services  
Helm deployment to k8s added
Makefile for docker/deploy scripts
glide install
add config values for gocloud-cicd deploy scripts
Test docker images as part of CI/CD  
go get -d fails to download github.com/stretchr/testify/assert in test file -- why?  
Add --log-level option cmds  
docker automated builds (suck): https://github.com/docker/hub-feedback/issues/1012  

# docker-compose
## start services
docker-compose up -d

## cleanup
docker-compose stop  
docker-compose rm -f

# kubernetes
## minikube
kubectl create -f kube/server-deploy.yaml  
kubectl create -f kube/server-service.yaml  
kubectl create -f kube/client-deploy.yaml  
kubectl expose deployment greeter-web --type=NodePort (minikube only!)  
kubectl get svc

NAME           CLUSTER-IP   EXTERNAL-IP   PORT(S)          AGE  
greeter-grpc   10.0.0.223   <none>        8000/TCP         3m  
greeter-web    10.0.0.58    <nodes>       8010:**PORT**/TCP   6m  

## testing
curl -i localhost:**PORT** ==> 404  
curl -i localhost:**PORT**/healthcheck ==> 200  
curl -i localhost:**PORT**/hw ==> Hello World! / 200  
curl -i localhost:**PORT**/hw?name=DUDE ==> Hello DUDE! / 200  

## cleanup
kubectl delete deploy/greeter-grpc deploy/greeter-web  
kubectl delete svc/greeter-grpc svc/greeter-web  
minikube stop

#NOTES
