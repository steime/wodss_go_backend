package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"strconv"
)

func (r *MySqlRepository) SaveAllModules(modules []persistence.Module) {
	for _, m := range modules {
		fmt.Println(m)
		r.db.Create(&m)
	}
}

func (r *MySqlRepository) GetAllModules() ([]persistence.Module,error){
	var modules []persistence.Module
	if result := r.db.Preload("Requirements").Find(&modules); result.Error != nil {
		return modules,result.Error
	}
	return modules,nil
}

func (r *MySqlRepository) GetModuleById(id string) (persistence.Module,error){
	var module persistence.Module
	i, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	if result := r.db.Preload("Requirements").Find(&module,i); result.Error != nil {
		return module,result.Error
	}
	return module,nil
}
