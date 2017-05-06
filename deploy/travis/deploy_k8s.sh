#!/usr/bin/env bash
# TODO: change positional args to flags ...OR... create python version anticipating reading in workflow JSON config

echo "helm/k8s install"
echo "args: $@"

# TODO: move to travis config and use as arg to script
# if [[ $TRAVIS_BRANCH =~ $BRANCH_REGEX ]];
# then export DOCKER_REPO=gcr.io/$GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME;
# else export DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]');
# fi

if [[ $# -ne 5 ]]
then
  echo "error: incorrect number of required positional args"
  echo "usage: deploy_k8s.sh repository tag release-name dry-run namespace"
  exit 1
fi

DOCKER_REPO=$1
COMMIT_TAG=$2
RELEASE_NAME=$3
DRYRUN=$4
NAMESPACE=$5

if [ $DRYRUN == "DRYRUN" ];
  then
    DRYRUN_OPTION=" --dry-run "
    echo "using --dry-run option; service not deployed."
fi

echo registry/repo: $DOCKER_REPO
echo commit tag: $COMMIT_TAG
echo release name: $RELEASE_NAME
echo dryrun: $DRYRUN_OPTION
echo namespace: $NAMESPACE
echo image: $DOCKER_REPO:COMMIT_TAG

# BUG: helm upgrade` does not re-create namespace if it's been deleted. https://github.com/kubernetes/helm/issues/2013
# create namespace all cases ignoring error
sudo kubectl get namespace $NAMESPACE || true

# upstall helm release
sudo helm upgrade \
$DRYRUN_OPTION \
--debug \
--install $RELEASE_NAME \
--namespace=$NAMESPACE \
--set service.gocloudAPI.image.repository=$DOCKER_REPO \
--set service.gocloudAPI.image.tag=$COMMIT_TAG \
--set service.gocloudGrpc.image.repository=$DOCKER_REPO \
--set service.gocloudGrpc.image.tag=$COMMIT_TAG \
deploy/helm/gocloud
