package main

import (
	"fmt"
	"strconv"
)

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
	Tag(string, string, string, int) ([]string, error)
	Push() (string, error)
	IsRegistryValid() bool
}

type GCRRegistry struct {
	Name    string
	Host    string
	Project string
	Repo    string
	Url     string
}

func (r *GCRRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (r *DockerRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (gcr *GCRRegistry) Tag(tag string, event string, branch string) ([]string, error) {
	var images []string
	image := gcr.Url + ":" + tag
	images = append(images, image)
	return images, nil
}

func (gcr *GCRRegistry) Push() (string, error) {
	return fmt.Sprintf("gcloud docker --push: %v\n", gcr.Url), nil
}

func (docker *DockerRegistry) Tag(tag string, event string, branch string, pr int) ([]string, error) {
	var images []string
	// build tag image
	image := docker.Url + ":" + tag
	images = append(images, image)

	fmt.Println("DEBUG Tag: ", tag, event, branch)
	switch event {
	case "push":
		images = append(images, docker.Url+":"+branch)
		if branch == "master" {
			images = append(images, docker.Url+":latest")
		}
	case "pull_request":
		images = append(images, docker.Url+":PR-"+strconv.Itoa(pr))
	}

	return images, nil
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
