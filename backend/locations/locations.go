package locations

import (
	"log"

	buildings_v1 "github.com/elonsoc/center/locations/v1/buildings"
	"github.com/go-chi/chi/v5"
)

// LocationsRouter is the router for the locations service
type LocationsRouter struct {
	chi.Router
	Logger *log.Logger
}

// Initialize locations
func NewLocationsRouter(l *LocationsRouter) *LocationsRouter {
	r := chi.NewRouter()
	l.Logger.SetPrefix("Locations: ")

	l.Logger.Println("Initializing locations service")

	// initialize the various endpoints
	r.Mount("/v1/buildings", buildings_v1.NewBuildingsRouter(&buildings_v1.BuildingsRouter{Logger: l.Logger}).Router)
	return &LocationsRouter{
		Router: r,
	}
}
