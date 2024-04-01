package buildings_v1

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/elonsoc/ods/backend/service"
	chi "github.com/go-chi/chi/v5"
)

// BuildingsRouter is the router for the buildings service
// It contains a pointer to the chi router that is defined in NewBuildingsRouter
// and a pointer to the service struct which is defined in the service package.
type BuildingsRouter struct {
	chi.Router
	Svcs *service.Services
}

// NewBuildingsRouter creates a new instance of the BuildingsRouter struct
// and returns a pointer to it.
// As defined in the NewLocationsRouter function in the locations.go file,
// this function takes a pointer to the BuildingsRouter struct as an argument
// so that we can pass the service struct around by reference after adding the handlers
// to the router. The reason we are passed a pointer to the BuildingsRouter struct
// is so that we can use the service struct in the top level to access the logger and other services
// that were defined in the service package.
func NewBuildingsRouter(b *BuildingsRouter) *BuildingsRouter {
	// At this point, the BuildingsRouter struct has been created
	// but the chi router has not been defined.

	// We need to define the chi router here so that we can
	// add the handlers to it.
	// Also, the service package is available to us here because
	// we passed a pointer to it when defining the BuildingsRouter struct as an argument
	r := chi.NewRouter()
	// We bind the handler functions to the router here.
	r.Get("/", b.RootHandler)
	r.Get("/{buildingID}", b.BuildingByIdHandler)
	// This router was not defined when we created the BuildingsRouter struct
	// in the parent function, so we need to set it here.
	// if we didn't, we would get a nil pointer error when we tried to use the router.
	b.Router = r
	return b
}

// RootHandler is the handler for the REST endpoint
// that will be used to get the buildings data.
// The endpoint for this handler is locations/v1/buildings
// As of right now, This endpoint returns a list of
// all the buildings and all of their attributes.
//
// It might be interesting to add a query parameter to this endpoint
// for filtering the buildings by building type.
func (br *BuildingsRouter) RootHandler(w http.ResponseWriter, r *http.Request) {
	buildings, err := br.Svcs.Db.GetBuildings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		br.Svcs.Log.Error("failed to get buildings: " + err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(buildings)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		br.Svcs.Log.Error("failed to encode buildings: " + err.Error(), nil)
	}
}

// BuildingByIdHandler is the handler for the REST endpoint
// that will be used to get a specific building's data.
// The endpoint for this handler is locations/v1/buildings/{buildingID}
// where buildingID is the id of the building that you want to get the data for.
func (br *BuildingsRouter) BuildingByIdHandler(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    w.Header().Set("Content-Type", "application/json")

    buildingId := strings.ToLower(chi.URLParam(r, "buildingID"))

    building, err := br.Svcs.Db.GetBuildingById(buildingId)
    if err != nil {
        br.Svcs.Log.Error("Building not found: "+buildingId, nil)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    br.Svcs.Stat.TimeElapsed("by_id.time", time.Since(start).Milliseconds())
    br.Svcs.Stat.Increment("by_id.count")

    err = json.NewEncoder(w).Encode(building)
    if err != nil {
        br.Svcs.Log.Error("Failed to encode building: "+err.Error(), nil)
        w.WriteHeader(http.StatusInternalServerError)
    }
}
