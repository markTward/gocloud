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
	IsRegistryValid() error
	Push([]string) (string, error)
	Authenticate() error
}

type GCRRegistry struct {
	Name        string
	Description string
	Host        string
	Project     string
	Repo        string
	Url         string
}

func (r *GCRRegistry) Authenticate() (err error) {

	var stderr bytes.Buffer
	cmd := exec.Command("gcloud", "auth", "activate-service-account", "--key-file", "client-secret.json")
	cmd.Stderr = &stderr

	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("%v", stderr.String())
	}
	return err

}

func (gcr *GCRRegistry) Push(images []string) (msg string, err error) {
	// TODO: real push!
	msg = fmt.Sprintf("gcloud docker --push %v", images)
	return msg, err
}

func (r *GCRRegistry) IsRegistryValid() (err error) {
	if r.Url == "" {
		err = fmt.Errorf("error: url missing from %v configuration", r.Description)
	}
	return err
}

type DockerRegistry struct {
	Name        string
	Description string
	Host        string
	Account     string
	Repo        string
	Url         string
}

func (r *DockerRegistry) Authenticate() (err error) {
	return err
}

func (r *DockerRegistry) IsRegistryValid() (err error) {
	if r.Url == "" {
		err = fmt.Errorf("error: url missing from %v configuration", r.Description)
	}
	return err
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
