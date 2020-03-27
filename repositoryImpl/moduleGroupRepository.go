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

func (r *MySqlRepository) GetAllModuleGroups() ([]persistence.ModuleGroup,error){
	var moduleGroups []persistence.ModuleGroup
	if result := r.db.Preload("Parent").Find(&moduleGroups); result.Error != nil {
		return moduleGroups,result.Error
	}
	if result := r.db.Preload("ModulesList").Find(&moduleGroups); result.Error != nil {
		return moduleGroups,result.Error
	}
	return moduleGroups,nil
}

func (r *MySqlRepository) GetModuleGroupById(id string) (persistence.ModuleGroup,error) {
	var moduleGroup persistence.ModuleGroup
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	if result := r.db.Preload("Parent").Find(&moduleGroup,i); result.Error != nil {
		return moduleGroup,result.Error
	}
	if result := r.db.Preload("ModulesList").Find(&moduleGroup,i); result.Error != nil {
		return moduleGroup,result.Error
	}
	return moduleGroup,nil
}