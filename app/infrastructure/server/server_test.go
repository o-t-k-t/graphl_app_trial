package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
)

func TestGraphQLEndpoint_ByTestData(t *testing.T) {
	// ARRANGE
	const testFileDir = "./testdata/queries"
	const expectFileDir = "./testdata/expected"

	// Get golden file list.
	files, err := os.ReadDir(testFileDir)
	if err != nil {
		t.Fatal(err)
	}

	// Build a server
	handler := createGraphQLHandler()
	defer closeGraphQLHandler()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		t.Run(file.Name(), func(t *testing.T) {
			// Read test case files
			query, err := ioutil.ReadFile(filepath.Join(testFileDir, file.Name()))
			assert.Equal(t, nil, err)

			// GQL Query to JSON
			jsonReq, err := json.Marshal(struct {
				Query string `json:"query"`
			}{
				Query: string(query),
			})
			assert.Equal(t, nil, err)

			// Create request from query file.
			req, err := http.NewRequestWithContext(
				context.Background(),
				http.MethodPost, "/query",
				bytes.NewReader(jsonReq),
			)
			assert.Nil(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Create response recoder.
			recorder := httptest.NewRecorder()

			// ACT
			handler.ServeHTTP(recorder, req)

			// ASSERT
			data := struct{ Type string }{Type: "Golden"}

			// compare recorded file. (or record file)
			goldie.New(t).AssertWithTemplate(
				t,
				file.Name(),
				data,
				recorder.Body.Bytes(),
			)
		})
	}
}
