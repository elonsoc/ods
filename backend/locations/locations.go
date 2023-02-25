package locations

import (
	buildings_v1 "github.com/elonsoc/center/locations/v1/buildings"
	"github.com/elonsoc/center/service"
	"github.com/go-chi/chi/v5"
)

// LocationsRouter is the router for the locations service
type LocationsRouter struct {
	chi.Router
	Svcs *service.Service
}

// Initialize locations
func NewLocationsRouter(l *LocationsRouter) *LocationsRouter {
	r := chi.NewRouter()
	l.Svcs.Logger.Println("Initializing locations service")

	// initialize the various endpoints
	r.Mount("/v1/buildings", buildings_v1.NewBuildingsRouter(&buildings_v1.BuildingsRouter{Svcs: l.Svcs}).Router)
	return &LocationsRouter{
		Router: r,
	}
}
