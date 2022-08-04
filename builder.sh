#!/usr/bin/env bash

CHART_DIR=${1:-$PWD}
REPOS=$(yq '.dependencies[] | .name + " " + .repository' $CHART_DIR/Chart.yaml)

while read -r REPO
do
 $HELM_BIN repo add $REPO
done < <(echo "$REPOS")

$HELM_BIN dependencies build $CHART_DIR
