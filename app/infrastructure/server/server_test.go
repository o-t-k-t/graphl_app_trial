package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGraphQLEndpoint_ByTestData(t *testing.T) {
	// Arange (server)
	ts := httptest.NewServer(createGraphQLHandler())
	defer ts.Close()

	// Arange (HTTP Request)
	q := struct{ Query string }{
		Query: "{listUsers {name}}",
	}
	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(&q); err != nil {
		t.Fatal("error encode", err)
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, ts.URL, &body)
	if err != nil {
		t.Fatal("error new request", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Act
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("error request", err)
	}
	defer res.Body.Close()

	// Assert
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("error read body", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("error request code:", res.StatusCode, string(result))
	}

	t.Log(string(result))
}
