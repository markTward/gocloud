package main

import "fmt"

type RestAPI struct {
	Name string
}
type RestAPIIniter interface {
	Init(string) error
}

func InitAPI(init RestAPIIniter, name string) error {
	fmt.Println("InitAPI: RestAPIIniter:", init)
	return init.Init(name)
}

func (api *RestAPI) Init(name string) error {
	api.Name = name
	return nil
}

func main() {
	api := &RestAPI{}
	InitAPI(api, "GoCloud")
	fmt.Printf("%#v", api)
}
