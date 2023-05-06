package buildings_v1

// Room is a struct that describes a room on a specific floor of a building.
// It contains the name of the room and the level of the floor that the room is on.
type Room struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Floor is a struct that describes a floor of a building.
// It contains the name of the floor, the level of the floor, the rooms on the floor,
// and the floorplan of the floor.
type Floor struct {
	Name      string `json:"name"`
	Level     int    `json:"level"`
	Rooms     []Room `json:"rooms"`
	Floorplan string `json:"floorplan"`
}

// LatLng is a struct that describes a location on the earth.
// It contains the latitude and longitude of the location.
type LatLng struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// BuildingType is an enum that describes the type of a building.
type BuildingType int64

// The following constants are the possible values of the BuildingType enum.
const (
	BuildingTypeUnknown   BuildingType = iota // Unknown Building Type
	BuildingTypeResidence                     // Residence Hall
	BuildingTypeDining                        // Dining Hall
	BuildingTypeOffice                        // Office Building
	BuildingTypeRetail                        // Retail Building
	BuildingTypeAcademic                      // Academic Building
	BuildingTypeOther                         // Other Building, not listed above
)

// String returns the string representation of a BuildingType.
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
	case BuildingTypeAcademic:
		return "Academic"
	case BuildingTypeOther:
		return "Other"
	default:
		return "Unknown"
	}
}

// Building is a struct that describes a building.
// It contains the name of the building, the floors of the building,
// the location of the building, the address of the building,
// the type of the building, and the id of the building.
type EnhancedBuilding struct {
	Name         string       `json:"name"`
	Floors       []Floor      `json:"floors"`
	Location     LatLng       `json:"location"`
	Address      string       `json:"address"`
	BuildingType BuildingType `json:"type"`
	Id           string       `json:"id"`
}
