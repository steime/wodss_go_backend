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

func (r *MySqlRepository) GetAllModules() []persistence.Module{
	var modules []persistence.Module
	r.db.Preload("Requirements").Find(&modules)
	return modules
}

func (r *MySqlRepository) GetModuleById(id string) persistence.Module{
	var module persistence.Module
	i, err := strconv.Atoi(id)
	r.db.Preload("Requirements").Find(&module,i)
	if err != nil {
		panic(err)
	}

	return module
}