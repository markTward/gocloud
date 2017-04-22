#!/usr/bin/env bash
# TODO: handle pull request
echo "docker_push.sh script start"
docker version

if [[ $TRAVIS_BRANCH =~ $BRANCH_REGEX ]];
then
  DOCKER_REPO=gcr.io/$GCLOUD_PROJECT_ID/$GOCLOUD_PROJECT_NAME;
  REPO_TARGET=gcr
else
  DOCKER_REPO=$(echo $TRAVIS_REPO_SLUG | tr '[:upper:]' '[:lower:]');
  REPO_TARGET=docker
fi

echo "DOCKER_REPO: $DOCKER_REPO"
echo "REPO_TARGET: $REPO_TARGET"
echo "DOCKER_COMMIT_TAG: $DOCKER_COMMIT_TAG"
echo "BRANCH_REGEX: $BRANCH_REGEX"
echo "TRAVIS_BRANCH: $TRAVIS_BRANCH"
echo "TRAVIS_EVENT_TYPE: $TRAVIS_EVENT_TYPE"

# tag images
docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:$DOCKER_COMMIT_TAG;

# pull requests

if [ "$TRAVIS_EVENT_TYPE" == "push" ]; then
  docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:$TRAVIS_BRANCH;
  if [ "$TRAVIS_BRANCH" == "master" ]; then
    docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:latest;
  fi
elif [ "$TRAVIS_EVENT_TYPE" == "pull_request" ]; then
  docker tag $GOCLOUD_PROJECT_NAME:$DOCKER_COMMIT_TAG $DOCKER_REPO:PR-$TRAVIS_PULL_REQUEST;
else
  echo "can't assign repo, unknown TRAVIS_EVENT_TYPE: $TRAVIS_EVENT_TYPE"
  exit 1
fi

docker images

# push images
if [ "$REPO_TARGET" == "gcr" ]; then
  if [ "$TRAVIS_EVENT_TYPE" == "push" ]; then
    sudo gcloud docker -- push $DOCKER_REPO:$DOCKER_COMMIT_TAG
    sudo gcloud docker -- push $DOCKER_REPO:$TRAVIS_BRANCH
    if [ "$TRAVIS_BRANCH" == "master" ]; then
      sudo gcloud docker -- push $DOCKER_REPO:latest
    fi
  elif [ "$TRAVIS_EVENT_TYPE" == "pull_request" ]; then
    sudo gcloud docker -- push $DOCKER_REPO:PR-$TRAVIS_PULL_REQUEST
  else
    echo "can't assign repo, unknown TRAVIS_EVENT_TYPE: $TRAVIS_EVENT_TYPE"
    exit 1
  fi
elif [ "$REPO_TARGET" == "docker" ]; then
  docker login -e $DOCKER_EMAIL -u $DOCKER_USER -p $DOCKER_PASSWORD
  docker push $DOCKER_REPO
else
  echo "can't assign target repo"
  exit 1
fi
