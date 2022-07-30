package main

import (
	"os"
	"path"
	"testing"
)

func TestGoDependenci(t *testing.T) {
	chartYaml, err := os.ReadFile(path.Join("testFiles", "Chart.yaml"))

	handleError(err)

	repos := getDependencies(chartYaml)

	if len(repos) <= 2 {
		t.Error("Did not find all Repositories")
	}

	if repos[0].Name != "kube-state-metrics" {
		t.Error("Failed to get repo name")
	}

	if repos[0].Repository != "https://prometheus-community.github.io/helm-charts" {
		t.Error("Failed to get repo url")
	}
}
