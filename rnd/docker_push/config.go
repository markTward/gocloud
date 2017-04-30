package main

import (
	"bytes"
	"fmt"
	"os/exec"
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
	IsRegistryValid() bool
	Push([]string) (string, error)
	Authenticate() error
}

type GCRRegistry struct {
	Name    string
	Host    string
	Project string
	Repo    string
	Url     string
}

type DockerRegistry struct {
	Name    string
	Host    string
	Account string
	Repo    string
	Url     string
}

func (r *GCRRegistry) Authenticate() (err error) {
	var stderr bytes.Buffer

	basecmd := "gcloud"
	args := []string{"auth", "activate-service-account", "--key-file", "client-secret.json"}
	cmd := exec.Command(basecmd, args...)
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("%v\n", stderr.String())
	}
	return err
}

func (r *DockerRegistry) Authenticate() (err error) {
	return err
}

func (r *GCRRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (r *DockerRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (gcr *GCRRegistry) Push(images []string) (msg string, err error) {
	if err = gcr.Authenticate(); err == nil {
		// TODO: real push!
		msg = fmt.Sprintf("gcloud docker --push %v\n", images)
	}
	return msg, err
}

func (docker *DockerRegistry) Push(images []string) (msg string, err error) {
	if err = docker.Authenticate(); err == nil {
		// TODO: real push!
		msg = fmt.Sprintf("docker push %v\n", images)
	}
	return msg, err
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
