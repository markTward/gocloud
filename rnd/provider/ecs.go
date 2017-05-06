package main

import (
	"fmt"
	"log"
)

type ECS struct {
	name         string
	awsAccessKey string
	awsSecretKey string
}

func (ecs *ECS) Install(svc string, creds string) (string, error) {
	fmt.Println(ecs, svc, creds)
	return "ECS install test", nil
}

func (ecs *ECS) Upgrade(svc string) (string, error) {
	fmt.Println(ecs, svc)
	return "ECS upgrade test", nil
}

func (ecs *ECS) Authenticate() (string, error) {
	// func (ecs *ECS) Authenticate(i ...interface{}) (string, error) {
	msg, err := ecs.getCredentials(ecs.awsAccessKey, ecs.awsSecretKey)
	return msg, err
}

func (ecs *ECS) getCredentials(key, secret string) (string, error) {
	var token string
	var err error

	if len(key) == 0 || len(secret) == 0 {
		return "", fmt.Errorf("missing aws key or secret: %v / %v", len(key) == 0, len(secret) == 0)
	}

	log.Println("attempt authenticate AWS creds")
	token, err = "dG9rZW4gcGxhY2Vob2xkZXIK", nil

	return token, err
}
