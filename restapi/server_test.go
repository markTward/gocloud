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
		url                string
		expectedStatusCode int
		message            string
	}{
		{
			description:        "successful query",
			url:                "/hw?name=DUDE",
			expectedStatusCode: 200,
			message:            "Hello DUDE!",
		},
		{
			description:        "successful query",
			url:                "/hw?name=GoCloud&name=DUDE",
			expectedStatusCode: 200,
			message:            "Hello GoCloud, DUDE!",
		},
		{
			description:        "successful query",
			url:                "/hw",
			expectedStatusCode: 200,
			message:            "Hello World!",
		},
	}

	// setup test RestAPI using local endpoint
	api := &RestAPI{}
	api.HelloWorlder = &TestHelloWorldEndpoint{}

	for _, hw := range tests {
		req, err := http.NewRequest("GET", hw.url, nil)
		if err != nil {
			t.Error(err)
		}

		w := httptest.NewRecorder()
		api.HelloWorldHandler(w, req)

		if hw.expectedStatusCode != w.Code {
			t.Errorf("handler returned wrong status code: got %v want %v",
				w.Code, hw.expectedStatusCode)
		}

		assert.Equal(t, hw.message, w.Body.String(), hw.description)

	}
}
