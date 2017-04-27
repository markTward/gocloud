package main

import "fmt"

type Provider interface {
	Authenticator
	Installer
	Upgrader
}

type Authenticator interface {
	Authenticate() (string, error)
	// Authenticate(...interface{}) (string, error)
}

type Upgrader interface {
	Upgrade(string) (string, error)
}

type Installer interface {
	Install(string, string) (string, error)
}

func install(p Provider, svc string) (string, error) {
	creds, err := p.Authenticate()
	if err != nil {
		return "", err
	}
	return p.Install(svc, creds)
}

func main() {
	p1 := &GKE{
		name:    "GKE",
		project: "k8sdemo",
		cluster: "k0",
	}

	fmt.Printf("%#v\n", p1)

	release, err := install(p1, "helloWorld")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("install release:", release)
	}

	fmt.Println()

	p2 := &ECS{
		name:         "ECS",
		awsAccessKey: "alkdjfladjsfk",
		awsSecretKey: "aldjfakdjklfkadjslfk",
	}

	fmt.Printf("%#v\n", p2)

	release, err = install(p2, "helloWorld micro service")
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("install release:", release)
	}
}
