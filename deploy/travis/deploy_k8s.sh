#!/usr/bin/env bash
echo "helm install start"
sudo helm list --all
sudo helm upgrade --dry-run --debug --install gocloud --namespace=gocloud --set service.gocloudAPI.image.tag=$DOCKER_COMMIT_TAG --set service.gocloudGrpc.image.tag=$DOCKER_COMMIT_TAG deploy/helm/gocloud/
sudo helm list --all

sudo kubectl get svc,deploy,pod --namespace=gocloud
