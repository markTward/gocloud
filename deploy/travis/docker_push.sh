#!/usr/bin/env bash
# TODO: handle pull request
echo "docker_push.sh script start"
docker version

# if [[ $TRAVIS_BRANCH =~ $BRANCH_REGEX ]];
# then
#   DOCKER_REPO=gcr.io/$GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME;
#   REPO_TARGET=gcr
# else
#   DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]');
#   REPO_TARGET=docker
# fi

DOCKER_REPO=gcr.io/$GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME

echo "DOCKER_REPO: $DOCKER_REPO"
echo "DOCKER_COMMIT_TAG: $DOCKER_COMMIT_TAG"
echo "TRAVIS_EVENT_TYPE=$TRAVIS_EVENT_TYPE"
echo "BRANCH_REGEX: $BRANCH_REGEX"
echo "TRAVIS_BRANCH: $TRAVIS_BRANCH"
echo "TRAVIS_EVENT_TYPE: $TRAVIS_EVENT_TYPE"

# tag images
docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:$DOCKER_COMMIT_TAG;
docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:$TRAVIS_BRANCH;
if [ "$TRAVIS_BRANCH" == "master" ]; then
  docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:latest;
fi
docker images

if [ $REPO_TARGET == "gcr"]; then
  sudo gcloud docker -- push $DOCKER_REPO:$DOCKER_COMMIT_TAG
  sudo gcloud docker -- push $DOCKER_REPO:$TRAVIS_BRANCH
  if [ "$TRAVIS_BRANCH" == "master" ]; then
    sudo gcloud docker -- push $DOCKER_REPO:latest
  fi
elif [ $REPO_TARGET == "docker"]; then
  docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASSWORD
  docker push $DOCKER_REPO
else
  echo "can't assign target repo"
fi
