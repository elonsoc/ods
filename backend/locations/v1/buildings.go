package buildings_v1

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

var logger = log.New(os.Stdout, "locations/v1/buildings: ", log.LstdFlags)

type Room struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type Floor struct {
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Rooms     []Room `json:"rooms"`
	Floorplan string `json:"floorplan"`
}

type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type BuildingType int64

const (
	BuildingTypeUnknown BuildingType = iota
	BuildingTypeResidence
	BuildingTypeDining
	BuildingTypeOffice
	BuildingTypeRetail
	BuildingTypeSchool
	BuildingTypeOther
)

func (t BuildingType) String() string {
	switch t {
	case BuildingTypeUnknown:
		return "Unknown"
	case BuildingTypeResidence:
		return "Residence"
	case BuildingTypeDining:
		return "Dining"
	case BuildingTypeOffice:
		return "Office"
	case BuildingTypeRetail:
		return "Retail"
	case BuildingTypeSchool:
		return "School"
	case BuildingTypeOther:
		return "Other"
	default:
		return "Unknown"
	}
}

type Building struct {
	Name         string       `json:"name"`
	Floors       []Floor      `json:"floors"`
	Location     LatLng       `json:"location"`
	Address      string       `json:"address"`
	BuildingType BuildingType `json:"type"`
	Id           string       `json:"id"`
}

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

func Buildings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BUILDINGS)
}

func BuildingWithId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	buildingId := strings.ToLower(chi.URLParam(r, "buildingID"))
	if BUILDINGS[buildingId].Name == "" {
		logger.Println("Building not found:", buildingId)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(BUILDINGS[buildingId])
}
