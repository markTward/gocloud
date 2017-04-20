#!/usr/bin/env bash
# TODO: change positional args to flags ...OR... create python version anticipating reading in workflow JSON config

echo "helm/k8s install"
echo "args: $@"

# TODO: move to travis config and use as arg to script
if [[ $TRAVIS_BRANCH =~ $BRANCH_REGEX ]];
then export DOCKER_REPO=gcr.io/GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME;
else export DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]');
fi

if [[ $# -ne 3 ]]
then
  echo "error: incorrect number of required positional args"
  echo "usage: deploy_k8s.sh deploy-type dry-run namespace"
  exit 1
fi

DEPLOY_TYPE=$1
DRYRUN=$2
NAMESPACE=$3

if [ $DRYRUN == "DRYRUN" ];
  then
    DRYRUN_OPTION=" --dry-run "
    echo "using --dry-run option; service not deployed."
fi

echo deployment type: $DEPLOY_TYPE
echo dryrun: $DRYRUN_OPTION
echo namespace: $NAMESPACE
echo project: $GOCLOUD_PROJECT_NAME
echo branch: $TRAVIS_BRANCH
echo image: $DOCKER_REPO:$DOCKER_COMMIT_TAG

sudo helm upgrade \
$DRYRUN_OPTION \
--debug \
--install $GOCLOUD_PROJECT_NAME \
--namespace=$NAMESPACE \
--set service.gocloudAPI.image.repository=$DOCKER_REPO \
--set service.gocloudAPI.image.tag=$DOCKER_COMMIT_TAG \
--set service.gocloudGrpc.image.repository=$DOCKER_REPO \
--set service.gocloudGrpc.image.tag=$DOCKER_COMMIT_TAG \
deploy/helm/gocloud
