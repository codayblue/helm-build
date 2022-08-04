# Helm Builder

This is a simple helm plugin that will take the dependent repos for a chart and set them up and then call the dependency build to download the charts from the repos.

## Requirements

To use this software you need to have [yq](https://github.com/mikefarah/yq) installed at version `>=4.0`

## Installation

Just run the command below to install

```bash
helm plugin install https://github.com/codayblue/helm-build
```

## Usage

While in the same directory as the helm chart run the following command

```bash
helm build
```
