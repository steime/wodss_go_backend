package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"strconv"
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

func (r *MySqlRepository) GetModuleGroupById(id string) persistence.ModuleGroup {
	var moduleGroup persistence.ModuleGroup
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	r.db.Preload("Parent").Find(&moduleGroup,i)
	r.db.Preload("ModulesList").Find(&moduleGroup,i)
	return moduleGroup
}
