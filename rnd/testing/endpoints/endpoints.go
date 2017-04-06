package endpoints

type HelloWorldEndpoint interface {
	HelloWorld(string) string
}

type RestAPIHelloWorldEndpoint struct{}

func (api RestAPIHelloWorldEndpoint) HelloWorld(name string) string {
	return "Hello " + name
}
