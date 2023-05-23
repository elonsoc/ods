package applications_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elonsoc/ods/backend/applications"
	"github.com/elonsoc/ods/backend/mocks"
	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Setting up the test server, mock db, mock logger, and testRequest function
func setup(t *testing.T) (*httptest.Server, *mocks.DbIFace, *mocks.LoggerIFace,
	func(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, []byte)) {
	db := mocks.NewDbIFace(t)
	logger := mocks.NewLoggerIFace(t)

	r := chi.NewRouter()

	ts := httptest.NewServer(r)
	Applications := applications.NewApplicationsRouter(
		&applications.ApplicationsRouter{
			Svcs: &service.Services{
				Db:  db,
				Log: logger,
			},
		},
	)

	r.Mount("/applications", Applications.Router)

	testRequest := func(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, []byte) {
		req, err := http.NewRequest(method, ts.URL+path, body)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()
		return resp, respBody
	}
	return ts, db, logger, testRequest
}

// Testing the NewApp function
func TestNewApp(t *testing.T) {
	// Arrange
	ts, db, _, testRequest := setup(t)
	defer ts.Close()

	// creating faux form data
	mockBody := map[string]string{
		"name":        "Twitter",
		"description": "A social media platform",
		"owners":      "Elon Musk",
		"teamName":    "SpaceX",
	}

	// converting the faux form data to json then to a reader
	bodyJson, err := json.Marshal(mockBody)
	if err != nil {
		t.Fatal(err)
	}

	buffer := bytes.NewBuffer(bodyJson)

	bodyReader := io.Reader(buffer)

	// mocking the database
	db.Mock.On("CheckDuplicate", mock.Anything, mock.Anything).Return(true, nil)
	db.Mock.On("NewApp", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act - make the request
	resp, body := testRequest(t, ts, http.MethodPost, "/applications/", bodyReader)

	// Assert - check the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	db.AssertCalled(t, "NewApp", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	db.AssertCalled(t, "CheckDuplicate", mock.Anything, mock.Anything)
	assert.Empty(t, body)
}
