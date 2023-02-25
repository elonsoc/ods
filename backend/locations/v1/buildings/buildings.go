package buildings_v1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// this map represents a database
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

type BuildingsRouter struct {
	chi.Router
	Logger *log.Logger
}

func NewBuildingsRouter(b *BuildingsRouter) *BuildingsRouter {
	r := chi.NewRouter()
	// add routes
	r.Get("/", b.RootHandler)
	r.Get("/{buildingID}", b.BuildingByIdHandler)
	b.Logger.SetPrefix("BuildingsRouter: ")
	b.Router = r
	return b
}

func (svc *BuildingsRouter) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BUILDINGS)
}

func (svc *BuildingsRouter) BuildingByIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	buildingId := strings.ToLower(chi.URLParam(r, "buildingID"))
	if BUILDINGS[buildingId].Name == "" {
		svc.Logger.Println("Building not found:", buildingId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(BUILDINGS[buildingId])
}
