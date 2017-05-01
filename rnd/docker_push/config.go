package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
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
	Push([]string) ([]string, error)
	Authenticate() error
}

type GCRRegistry struct {
	Name        string
	Description string
	Host        string
	Project     string
	Repo        string
	Url         string
	KeyFile     string
}

func (r *GCRRegistry) getClientID() (email string, err error) {

	// parse google credentials for identity
	type clientSecret struct {
		ClientEmail string `json:"client_email"`
	}

	// read in project config file
	jsonInput, err := ioutil.ReadFile(r.KeyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	// parse yaml into Config object
	ce := clientSecret{}
	err = json.Unmarshal([]byte(jsonInput), &ce)
	log.Println(ce)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return ce.ClientEmail, nil
}

func (r *GCRRegistry) Authenticate() (err error) {
	var stderr bytes.Buffer

	cmd := exec.Command("gcloud", "auth", "activate-service-account", "--key-file", r.KeyFile)
	cmd.Stderr = &stderr

	cid, _ := r.getClientID()

	log.Printf("attempt gcr authenication: %v\n", cid)
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("%v", stderr.String())
	} else {
		log.Printf("gcr authenication successful: %v\n", cid)
	}

	return err

}

func (gcr *GCRRegistry) Push(images []string) (pushed []string, err error) {
	var stderr bytes.Buffer
	var cmdOut []byte
	// IDEA: could use single command to push all repo images: gcloud docker -- push gcr.io/k8sdemo-159622/gocloud
	// but assumes that process ALWAYS wants ALL tags for repo to be pushed.  good for isolated build env, but ...
	for _, image := range images {
		log.Println("attempt push to gcr: ", image)

		cmd := exec.Command("gcloud", "docker", "--", "push", image)
		cmd.Stderr = &stderr

		if cmdOut, err = cmd.Output(); err != nil {
			err = fmt.Errorf("%v: %v", image, stderr.String())
			break
		}

		for _, o := range strings.Split(strings.TrimSpace(string(cmdOut)), "\n") {
			log.Println(o)
		}

		pushed = append(pushed, image)
	}
	return pushed, err
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
	var stderr bytes.Buffer

	dockerUser := os.Getenv("DOCKER_USER")
	if dockerUser == "" {
		err = fmt.Errorf("%v", "error: DOCKER_USER environment variable not set")
		log.Println(err)
		return err
	}

	dockerPass := os.Getenv("DOCKER_PASSWORD")
	if dockerPass == "" {
		err = fmt.Errorf("%v", "error: DOCKER_PASSWORD environment variable not set")
		log.Println(err)
		return err
	}

	cmd := exec.Command("docker", "login", "-u", dockerUser, "-p", dockerPass)
	cmd.Stderr = &stderr

	log.Printf("attempt docker login: %v\n", dockerUser)
	if err = cmd.Run(); err != nil {
		err = fmt.Errorf("%v", stderr.String())
		return err
	}

	log.Printf("docker login successful: %v\n", dockerUser)

	return err
}

func (r *DockerRegistry) IsRegistryValid() (err error) {
	if r.Url == "" {
		err = fmt.Errorf("error: url missing from %v configuration", r.Description)
	}
	return err
}

func (docker *DockerRegistry) Push(images []string) (pushed []string, err error) {
	var stderr bytes.Buffer
	var cmdOut []byte

	for _, image := range images {
		log.Println("attempt push to docker registry: ", image)

		cmd := exec.Command("docker", "push", image)
		cmd.Stderr = &stderr

		if cmdOut, err = cmd.Output(); err != nil {
			err = fmt.Errorf("%v: %v", image, stderr.String())
			break
		}

		for _, o := range strings.Split(strings.TrimSpace(string(cmdOut)), "\n") {
			log.Println(o)
		}

		pushed = append(pushed, image)
	}
	return pushed, err
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
