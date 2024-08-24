package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldShouldSucceed(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()
	testClient := testServer.Client()

	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "Hello World\n", string(body))
	assert.Equal(t, 200, response.StatusCode)

}

func TestHelloWorldShoulFail(t *testing.T) {
	//create a test server
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	// it is required that test server be closed
	defer testServer.Close()
	// create a test client to immitate user
	testClient := testServer.Client()
	// have the test client send a get request to URL created by test server
	body := strings.NewReader("some body")

	response, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 405, response.StatusCode)

}
