package persistence

type Profile struct {
	ID string `gorm:"number(7);PRIMARY_KEY;" json:"id"`
	Name string `gorm:"type:varchar(100)" json:"name"`
	ListOfModules []ListOfModules `json:"modules"`
	Minima uint `gorm:"number(3);" json:"minima"`
}

type ListOfModules struct {
	ProfileID string `json:"-"`
	ModuleID string `json:"id"`
}

type ProfileResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ListOfModules []string `json:"modules"`
	Minima uint `json:"minima"`
}
