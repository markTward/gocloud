#!/usr/bin/env bash
echo "helm install start"
echo "args: $@"

if [[ $TRAVIS_BRANCH =~ $BRANCH_REGEX ]];
then export DOCKER_REPO=gcr.io/GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME;
else export DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]');
fi

if [[ $# -eq 2 ]] && [ $2 == "DRYRUN" ];
  then
    DRYRUN_OPTION=" --dry-run "
    echo "using --dry-run option; service not deployed."
fi

sudo helm upgrade \
-- debug --$DRYRUN_OPTION \
--install $GOCLOUD_PROJECT_NAME \
--namespace=gocloud \
--set service.gocloudAPI.image.repository=$DOCKER_REPO \
--set service.gocloudAPI.image.tag=$DOCKER_COMMIT_TAG \
--set service.gocloudGrpc.image.repository=$DOCKER_REPO \
--set service.gocloudGrpc.image.tag=$DOCKER_COMMIT_TAG \
deploy/helm/gocloud
