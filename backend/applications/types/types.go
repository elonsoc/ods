package types

// The BaseApplication type defines the structure of an application at its infancy.
type BaseApplication struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Owners      string `json:"owners"`
}

type Application struct {
	BaseApplication
	Id      string `json:"id" db:"id"`
	ApiKey  string `json:"apiKey" db:"api_key"`
	IsValid bool   `json:"isValid" db:"is_valid"`
}
