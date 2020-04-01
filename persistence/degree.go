package persistence

type Degree struct {
	ID string `gorm:"number(7);PRIMARY_KEY;" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	Groups []Groups `json:"groups"`
}

type Groups struct {
	DegreeID string `json:"-"`
	GroupID string	`json:"id"`
}
