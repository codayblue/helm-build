package main

import (
	"os"
	"os/exec"
	"path"

	"gopkg.in/yaml.v3"
)

type Repo struct {
	Name       string
	Repository string
}

type Repositories struct {
	Dependencies []Repo
}

type HelmCommand struct {
	helmbin string
	stdout  *os.File
	stderr  *os.File
}

func (helmcommand *HelmCommand) execute(args ...string) {
	cmd := exec.Command(helmcommand.helmbin, args...)

	cmd.Stdout = helmcommand.stdout
	cmd.Stderr = helmcommand.stderr

	err := cmd.Run()

	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func getDependencies(chartYaml []byte) []Repo {
	var repos Repositories

	err := yaml.Unmarshal(chartYaml, &repos)

	handleError(err)

	return repos.Dependencies
}

func addRepos(helmcommand HelmCommand, repos []Repo) {
	for _, repo := range repos {
		helmcommand.execute("repo", "add", repo.Name, repo.Repository)
	}
}

func main() {
	currentDir, err := os.Getwd()

	handleError(err)

	chartYaml, err := os.ReadFile(path.Join(currentDir, "Chart.yaml"))

	handleError(err)

	repos := getDependencies(chartYaml)
	helmcommand := HelmCommand{
		helmbin: os.Getenv("HELM_BIN"),
		stdout:  os.Stdout,
		stderr:  os.Stderr,
	}

	addRepos(helmcommand, repos)

	helmcommand.execute("dependencies", "build")
}
