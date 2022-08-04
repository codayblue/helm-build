#!/usr/bin/env bash

REPOS=$(yq '.dependencies[] | .name + " " + .repository' $PWD/Chart.yaml)

while read -r REPO
do
 $HELM_BIN repo add $REPO
done < <(echo "$REPOS")

$HELM_BIN dependencies build