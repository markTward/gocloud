#!/usr/bin/env bash
echo "helm install start"
echo "args: $@"

sudo helm upgrade --dry-run --debug --install gocloud --namespace=gocloud \
--set service.gocloudAPI.image.repository=$DOCKER_REPO \
--set service.gocloudAPI.image.tag=$DOCKER_COMMIT_TAG \
--set service.gocloudGrpc.image.repository=$DOCKER_REPO \
--set service.gocloudGrpc.image.tag=$DOCKER_COMMIT_TAG \
deploy/helm/gocloud
