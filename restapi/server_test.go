package restapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHelloWorldEndpoint struct{}

func (api TestHelloWorldEndpoint) HelloWorld(names []string) (string, error) {
	// handle 0-to-Many qs names
	defaultName := "World"
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
		endpoint           *TestHelloWorldEndpoint
		url                string
		expectedStatusCode int
		message            string
	}{
		{
			description:        "successful query",
			endpoint:           &TestHelloWorldEndpoint{},
			url:                "/hw?name=DUDE",
			expectedStatusCode: 200,
			message:            "Hello DUDE!",
		},
		{
			description:        "successful query",
			endpoint:           &TestHelloWorldEndpoint{},
			url:                "/hw",
			expectedStatusCode: 200,
			message:            "Hello World!",
		},
	}

	for _, hw := range tests {
		app := &RestAPI{}
		app.HelloWorldEndpoint = hw.endpoint

		req, err := http.NewRequest("GET", hw.url, nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		app.HelloWorldHandler(w, req)

		if hw.expectedStatusCode != w.Code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				w.Code, hw.expectedStatusCode)
		}

		assert.Equal(t, hw.message, w.Body.String(), hw.description)

	}
}
