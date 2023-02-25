package locations

import (
	buildings_v1 "github.com/elonsoc/center/locations/v1"
	"github.com/go-chi/chi/v5"
)

// Initialize locations
func Initialize(router chi.Router) {
	// Initialize v1 the v1 route on the toplevel router
	router.Route("/locations/v1/", func(r chi.Router) {
		r.Route("/buildings", func(r chi.Router) {
			r.Get("/", buildings_v1.Buildings)
			r.Get("/{buildingID}", buildings_v1.BuildingWithId)
		})

	})
}
