#!/usr/bin/env bash
echo "helm install start"
echo "args: $@"
echo "arg count: $#"
echo "special arg: $0"

if [[ $# -eq 2 ]] && [ $2 == "DRYRUN" ];
  then
    echo "helm dry-run enabled; containers not deployed!"
    DRYRUN_OPTION="--dry-run"
fi

TRAVIS_BRANCH=$3
if [ $TRAVIS_BRANCH == "master" ]; then NAMESPACE=gocloud; elif [ $TRAVIS_BRANCH == "dev" ]; then NAMESPACE=gocloud-dev; elif [[ $TRAVIS_BRANCH =~ ^pr/.* ]]; then NAMESPACE=gocloud-stg; elif [[ $TRAVIS_BRANCH =~ ^(feature$|feature/.*|feature-.*) ]]; then NAMESPACE=; else echo "no branch $TRAVIS_BRANCH namespace match"; NAMESPACE=; fi

echo "$TRAVIS_BRANCH = $NAMESPACE"
