// Module Group Models, with JSON names and validation tags
package persistence

type ModuleGroup struct {
	ID string `gorm:"number(7);PRIMARY_KEY;" json:"id"`
	Name string `gorm:"type:varchar(100);" json:"name"`
	Minima uint `gorm:"number(3);" json:"minima"`
	Parent Parent `json:"parent"`
	ModulesList []ModulesList `json:"modules"`
}

type Parent struct {
	ModuleGroupID string `json:"-"`
	Parent *string `json:"id"`
}

type ModulesList struct {
	ModuleGroupID string `json:"-"`
	ModuleID string `json:"id"`
}
// For JSON ID Object List to String List conversion in Handler
type ModuleGroupsResponse struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Minima uint `json:"minima"`
	Parent *string `json:"parent"`
	ModulesList []string `json:"modules"`
}
