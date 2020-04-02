package mySQL

import "github.com/steime/wodss_go_backend/persistence"

func (r *MySqlRepository) CreateModuleVisit(visit *persistence.ModuleVisit) (*persistence.ModuleVisit,error) {
	if result := r.db.Create(&visit); result.Error != nil {
		return visit,result.Error
	} else {
		return visit,nil
	}
}
