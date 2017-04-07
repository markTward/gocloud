package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestHelloWorldEndpoint struct{}

func (api TestHelloWorldEndpoint) HelloWorld(names []string) (string, error) {
	// handle 0-to-Many qs names
	defaultName := "Hello World!"
	name := defaultName
	if len(names) != 0 {
		name = strings.Join(names, ", ")
	}

	msg := fmt.Sprintf("Hello %s!", name)
	return msg, nil
}

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		description        string
		hwEndpoint         *TestHelloWorldEndpoint
		url                string
		expectedStatusCode int
		message            string
	}{
		{
			description:        "successful query",
			hwEndpoint:         &TestHelloWorldEndpoint{},
			url:                "/hw?name=DUDE",
			expectedStatusCode: 200,
			message:            "Hello DUDE!",
		},
	}

	for _, hw := range tests {
		app := &RestAPI{hw: hw.hwEndpoint}
		req, err := http.NewRequest("GET", hw.url, nil)
		if err != nil {

		}
		w := httptest.NewRecorder()
		app.HelloWorldHandler(w, req)
		if hw.expectedStatusCode != w.Code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				w.Code, hw.expectedStatusCode)
		}
		if hw.message != w.Body.String() {
			t.Errorf("handler return wrong body: got %v want %v", w.Body.String(), hw.message)
		}
	}
}
