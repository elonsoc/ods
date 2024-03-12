package applications_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/elonsoc/ods/src/api/applications"
	"github.com/elonsoc/ods/src/common"
	"github.com/elonsoc/ods/src/common/types"
	"github.com/elonsoc/ods/src/mocks"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewApp(t *testing.T) {
	// Arrange - setup the test
	db := mocks.NewDbIFace(t)
	logger := mocks.NewLoggerIFace(t)
	stat := mocks.NewStatIFace(t)
	tok := mocks.NewTokenIFace(t)
	r := chi.NewRouter()

	ts := httptest.NewServer(r)
	apps := applications.NewApplicationsRouter(
		&applications.ApplicationsRouter{
			Svcs: &common.Services{
				Db:    db,
				Log:   logger,
				Stat:  stat,
				Token: tok,
			},
		},
	)
	r.Mount("/", apps)

	// Mock services being called in function
	tok.Mock.On("GetUidFromToken", mock.AnythingOfType("string")).Return("1", nil)
	db.Mock.On("NewApp", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return("12345", nil)

	// create the request body
	app := types.BaseApplication{}

	app.Description = "Test App"
	app.Name = "Test App"

	appJson, err := json.Marshal(app)
	reqBody := strings.NewReader(string(appJson))

	// create the request
	req, err := http.NewRequest("POST", ts.URL, reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req.AddCookie(&http.Cookie{Name: "ods_login_cookie_nomnom", Value: "123"})

	// Act - execute the test
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	// Assert - check the results
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}

	s := string(respBody)

	assert.Equal(t, "12345", s, nil)
	tok.AssertCalled(t, "GetUidFromToken", mock.AnythingOfType("string"))
	db.AssertCalled(t, "NewApp", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"))

}
