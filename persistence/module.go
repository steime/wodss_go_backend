package persistence

type Module struct {
	ID string `gorm:"number(7);PRIMARY_KEY;" json:"id"`
	Name string `gorm:"type:varchar(30);" json:"name"`
	Code string `gorm:"type:varchar(10);" json:"code"`
	Credits uint `gorm:"number(7);" json:"credits"`
	Hs bool `gorm:"number(1);" json:"hs"`
	Fs bool `gorm:"number(1);" json:"fs"`
	Requirements `json:"requirements"`
}

type Requirements[] struct {
	ModuleID string `json:"id"`
}


type ModuleRepository interface {
	GetAllModules() []Module
}

type DefaultModuleRepository struct {
	Modules []Module
}