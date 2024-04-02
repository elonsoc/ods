package buildings_v1_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	bldg "github.com/elonsoc/ods/backend/locations/v1/buildings"
	types "github.com/elonsoc/ods/backend/locations/v1/buildings/types"
	"github.com/elonsoc/ods/backend/mocks"
	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t *testing.T) (*httptest.Server, *mocks.DbIFace,
	*mocks.LoggerIFace, *mocks.StatIFace,
	func(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, []byte),
) {
	db := mocks.NewDbIFace(t)
	logger := mocks.NewLoggerIFace(t)
	stat := mocks.NewStatIFace(t)

	r := chi.NewRouter()

	ts := httptest.NewServer(r)
	BuildingsV1 := bldg.NewBuildingsRouter(
		&bldg.BuildingsRouter{
			Svcs: &service.Services{
				Db:   db,
				Log:  logger,
				Stat: stat,
			},
		},
	)
	r.Mount("/buildings", BuildingsV1.Router)

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
	return ts, db, logger, stat, testRequest
}

func TestBuildingByIdHandler(t *testing.T) {
	// Arrange - setup the test
	/*
		Although setup creates a mock for db and logger, we don't need to use them
		in this test so we can ignore them.
	*/
	ts, _, _, stat, testRequest := setup(t)
	defer ts.Close()

	stat.Mock.On("Increment", mock.Anything)
	stat.Mock.On("TimeElapsed", mock.Anything, mock.Anything)

	mcewenBuildingMock := types.Building{
		ID:                       strPtr("mcewen"),
		Description:              strPtr("McEwen Dining Hall"),
		Location:                 strPtr("1600 Amphitheatre Parkway"),
		LocationRepresentation:   strPtr("Mountain View, CA 94043"),
		Type:                     strPtr("Dining"),
		TypeRepresentation:       strPtr("Dining Hall"),
		LongDescription:          strPtr("McEwen Dining Hall is a popular dining location with multiple floors and a variety of food options."),
		City:                     strPtr("Mountain View"),
		State:                    strPtr("CA"),
		Zip:                      strPtr("94043"),
		Sector:                   strPtr("Education"),
		SectorRepresentation:     strPtr("University Campus"),
		Latitude:                 float64Ptr(37.422),
		Longitude:                float64Ptr(-122.084),
		AddDate:                  timePtr(time.Now()),
		ChangeDate:               timePtr(time.Now()),
	}
	
	// Act - make the request
	resp, body := testRequest(t, ts, http.MethodGet, "/buildings/mcewen", nil)

	// Assert - check the response

	respBuilding := types.Building{}
	err := json.Unmarshal(body, &respBuilding)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, mcewenBuildingMock, respBuilding)
}

func TestBuildingByIdHandlerNotFound(t *testing.T) {
	// Arrange - setup the test
	ts, db, log, stat, testRequest := setup(t)
	defer ts.Close()

	log.Mock.On("Error", mock.Anything, mock.Anything)

	// Act - make the request
	resp, body := testRequest(t, ts, http.MethodGet, "/buildings/nothere", nil)

	// Assert - check the response
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	log.AssertCalled(t, "Error", "Building not found: nothere", mock.Anything)
	// the following should, one day, be called.
	stat.AssertNotCalled(t, "Increment", mock.Anything)
	stat.AssertNotCalled(t, "TimeElapsed", mock.Anything)
	db.AssertNotCalled(t, "GetBuildingById", mock.Anything)
	assert.Empty(t, body)
}

func TestRootHandler(t *testing.T) {
	// Arrange
	ts, _, _, _, testRequest := setup(t)
	defer ts.Close()

	// we will, one day, make a call to the database to get all the buildings
	// so it is mocked here.
	// db.Mock.On("GetAllBuildings").Return([]buildings_v1.Building{
	// 	{
	// 		Name: "McEwen Dining Hall",
	// 		Floors: []buildings_v1.Floor{
	// 			{Name: "Floor 1", Level: 1, Rooms: []buildings_v1.Room{{Name: "Room 1", Level: 1}, {Name: "Room 2", Level: 1}}},
	// 			{Name: "Floor 2", Level: 2, Rooms: []buildings_v1.Room{{Name: "Room 3", Level: 2}, {Name: "Room 4", Level: 2}}},
	// 		},
	// 		Location:     buildings_v1.LatLng{Lat: 37.422, Lng: -122.084},
	// 		Address:      "1600 Amphitheatre Parkway, Mountain View, CA 94043",
	// 		BuildingType: buildings_v1.BuildingTypeDining,
	// 		Id:           "mcewen",
	// 	},
	// })

	// Act - make the request
	resp, body := testRequest(t, ts, http.MethodGet, "/buildings", nil)

	// Assert - check the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Contains(t, string(body), "mcewen")
}

func strPtr(s string) *string { return &s }
func float64Ptr(f float64) *float64 { return &f }
func timePtr(t time.Time) *time.Time { return &t }