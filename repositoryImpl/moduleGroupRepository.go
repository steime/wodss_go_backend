package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
)

func (r *MySqlRepository) SaveAllModuleGroups(moduleGroups []persistence.ModuleGroup) {
	for _, m := range moduleGroups {
		fmt.Println(m)
		r.db.Create(&m)
	}
}

func (r *MySqlRepository) GetAllModuleGroups() []persistence.ModuleGroup{
	var moduleGroups []persistence.ModuleGroup
	r.db.Preload("Parent").Find(&moduleGroups)
	r.db.Preload("ModulesList").Find(&moduleGroups)
	return moduleGroups
}
