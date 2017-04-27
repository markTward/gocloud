package main

import (
	"fmt"
	"log"
)

type GKE struct {
	name    string
	project string
	cluster string
}

func (gke *GKE) Install(svc string, creds string) (string, error) {
	var err error
	log.Println("attempt install svc:", svc)
	// release, err := fmt.Sprintf("service installed as release: %v\n", svc+"-A"), nil
	release, err := "", fmt.Errorf("error: force install fail")
	return release, err
}

func (gke *GKE) Upgrade(svc string) (string, error) {
	fmt.Println(gke, svc)
	return "GKE upgrade test", nil
}

func (gke *GKE) Authenticate() (string, error) {
	// func (gke *GKE) Authenticate(i ...interface{}) (string, error) {
	msg := gke.getCredentials(gke.name)
	return msg, nil
}

func (gke *GKE) getCredentials(acct string) string {
	token := "dG9rZW4gcGxhY2Vob2xkZXIK"
	return token
}
