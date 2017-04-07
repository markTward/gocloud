package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/markTward/gocloud/restapi/handlers"
)

func TestHealthCheckHandler(t *testing.T) {
	// adopted from: https://elithrar.github.io/article/testing-http-handlers-go/

	req, err := http.NewRequest("GET", "/healthcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
