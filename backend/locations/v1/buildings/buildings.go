package buildings_v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/elonsoc/center/backend/service"
	"github.com/go-chi/chi/v5"
)

// this map is a mock database of buildings
var BUILDINGS = map[string]Building{
	"mcewen": {
		Name: "McEwen Dining Hall",
		Floors: []Floor{
			{Name: "Floor 1", Level: 1, Rooms: []Room{{Name: "Room 1", Level: 1}, {Name: "Room 2", Level: 1}}},
			{Name: "Floor 2", Level: 2, Rooms: []Room{{Name: "Room 3", Level: 2}, {Name: "Room 4", Level: 2}}},
		},
		Location:     LatLng{Lat: 37.422, Lng: -122.084},
		Address:      "1600 Amphitheatre Parkway, Mountain View, CA 94043",
		BuildingType: BuildingTypeDining,
		Id:           "mcewen",
	},

	"powell": {
		Name: "Powell Building",
		Floors: []Floor{
			{Name: "Floor 1", Level: 1, Rooms: []Room{{Name: "Room 1", Level: 1}, {Name: "Room 2", Level: 1}}},
			{Name: "Floor 2", Level: 2, Rooms: []Room{{Name: "Room 3", Level: 2}, {Name: "Room 4", Level: 2}}},
			{Name: "Floor 3", Level: 3, Rooms: []Room{{Name: "Room 5", Level: 3}, {Name: "Room 6", Level: 3}}},
		},
		Location:     LatLng{Lat: 37.422, Lng: -122.084},
		Address:      "1600 Amphitheatre Parkway, Mountain View, CA 94043",
		BuildingType: BuildingTypeOffice,
		Id:           "powell",
	},
	"lodge-a": {
		Name: "The Lodge — Dormitory A",
		Floors: []Floor{
			{Name: "First Floor", Level: 0, Rooms: []Room{{Name: "101", Level: 0}}, Floorplan: "https://eloncdn.blob.core.windows.net/eu3/sites/789/2018/04/Lodge-Dormitory-A-First-Floor.pdf"},
		},
		Location:     LatLng{Lat: 37.422, Lng: -122.084},
		Address:      "1600 Amphitheatre Parkway, Mountain View, CA 94043",
		BuildingType: BuildingTypeResidence,
		Id:           "lodge-a",
	},
}

// BuildingsRouter is the router for the buildings service
// It contains a pointer to the chi router that is defined in NewBuildingsRouter
// and a pointer to the service struct which is defined in the service package.
type BuildingsRouter struct {
	chi.Router
	Svcs *service.Service
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
func (be *BuildingsRouter) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BUILDINGS)
}

// BuildingByIdHandler is the handler for the REST endpoint
// that will be used to get a specific building's data.
// The endpoint for this handler is locations/v1/buildings/{buildingID}
// where buildingID is the id of the building that you want to get the data for.
func (be *BuildingsRouter) BuildingByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	buildingId := strings.ToLower(chi.URLParam(r, "buildingID"))
	if BUILDINGS[buildingId].Name == "" {
		be.Svcs.Logger.Println("Building not found:", buildingId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(BUILDINGS[buildingId])
}
