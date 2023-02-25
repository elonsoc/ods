package buildings_v1

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
