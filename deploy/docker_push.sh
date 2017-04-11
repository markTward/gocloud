#!/bin/bash
set -x

echo "docker_push.sh script start"

env | grep TRAVIS | sort
docker version

DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]')

docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASSWORD

if [ "$TRAVIS_EVENT_TYPE" == "pull_request" ]; then
  echo "docker build/push for TRAVIS_EVENT_TYPE=pull_request"
  docker tag $DOCKER_REPO:$TRAVIS_COMMIT $DOCKER_REPO:PR-$TRAVIS_PULL_REQUEST;
  docker images
  docker push $DOCKER_REPO:PR-$TRAVIS_PULL_REQUEST
fi

if [ "$TRAVIS_EVENT_TYPE" == "push" ]; then
  echo "docker build/push for TRAVIS_EVENT_TYPE=push"
  docker tag $DOCKER_REPO:$TRAVIS_COMMIT $DOCKER_REPO:$TRAVIS_BRANCH;
  if [ "$TRAVIS_BRANCH" == "master"]; then
    docker tag $DOCKER_REPO:$TRAVIS_COMMIT $DOCKER_REPO:latest;
  fi
  docker images
  docker push $DOCKER_REPO:$TRAVIS_BRANCH
fi
