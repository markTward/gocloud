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
	Push() (string, error)
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
	args := []string{"auth", "activate-service-account", "--key-file", "client-secret.jsonXXX"}
	cmd := exec.Command(basecmd, args...)
	cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("%v\n", stderr.String())
	}
	return err
}

func (r *DockerRegistry) Authenticate() bool {
	return true
}

func (r *GCRRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (r *DockerRegistry) IsRegistryValid() bool {
	return r.Url != ""
}

func (gcr *GCRRegistry) Push() (msg string, err error) {
	if err = gcr.Authenticate(); err == nil {
		// TODO: real push!
		msg = fmt.Sprintf("gcloud docker --push: %v\n", gcr.Url)
	}
	return msg, err
}

func (docker *DockerRegistry) Push() (string, error) {
	return fmt.Sprintf("docker push %v\n", docker.Url), nil
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
