package main

import "fmt"

type App struct {
	Name string
}
type Github struct {
	Repo string
}

type Registry struct {
	GCRRegistry
	DockerRegistry
}

type Registrator interface {
	Tag(string) (string, error)
	Push() (string, error)
}

type GCRRegistry struct {
	Name    string
	Host    string
	Project string
	Repo    string
	Url     string
}

func (gcr *GCRRegistry) Tag(tag string) (string, error) {
	return fmt.Sprintf("GCR Tagging with: %v\n", tag), nil
}

func (gcr *GCRRegistry) Push() (string, error) {
	return fmt.Sprintf("gcloud docker --p ush: %v\n", gcr.Url), nil
}

func (docker *DockerRegistry) Tag(tag string) (string, error) {
	return fmt.Sprintf("Docker Tagging with: %v\n", tag), nil
}

func (docker *DockerRegistry) Push() (string, error) {
	return fmt.Sprintf("docker push %v\n", docker.Url), nil
}

type DockerRegistry struct {
	Name    string
	Host    string
	Account string
	Repo    string
	Url     string
}

type Workflow struct {
	Enabled bool

	Github struct {
		Repo   string
		Branch string
	}

	CIProvider struct {
		Name string
		Plan string
	}

	Platform struct {
		Name    string
		Project string
		Cluster string
	}

	Registry string

	CDProvider struct {
		Name      string
		Release   string
		Namespace string
		ChartDir  string
	}
}
