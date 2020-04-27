// Degree Models, with JSON names
package persistence

type Degree struct {
	ID string `gorm:"number(7);PRIMARY_KEY;" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Groups []Groups `json:"groups"`
	ProfilesByDegree []ProfilesByDegree `json:"profiles"`
}

type Groups struct {
	DegreeID string `json:"-"`
	GroupID string	`json:"id"`
}

type ProfilesByDegree struct {
	DegreeID string		`json:"-"`
	ProfileID string	`json:"id"`
}
// For JSON ID Object List to String List conversion in Handler
type DegreeResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Groups []string `json:"groups"`
	Profiles []string `json:"profiles"`
}
