package mySQL

import (
	"errors"
	"github.com/steime/wodss_go_backend/persistence"
	"strconv"
)

func (r *MySqlRepository) CreateModuleVisit(visit *persistence.ModuleVisit) (*persistence.ModuleVisit,error) {
	studentID := visit.Student
	moduleID := visit.Module
	if !r.CheckIfStudentExists(studentID) {
		return visit,errors.New("student not existing")
	} else if !r.CheckIfModuleExists(moduleID) {
		return visit,errors.New("module not existing")
	} else if r.CheckIfModuleVisitExists(moduleID) {
		return visit,errors.New("moduleVisit for this module exists")
	} else {
		if result := r.db.Create(&visit); result.Error != nil {
			return visit,result.Error
		} else {
			return visit,nil
		}
	}
}

func (r *MySqlRepository) CheckIfStudentExists(id uint) bool {
	var student persistence.Student
	if result := r.db.Find(&student,id); result.Error != nil {
		return false
	} else {
		return true
	}
}

func (r *MySqlRepository) CheckIfModuleExists(id string) bool {
	var module persistence.Module
	if i, err := strconv.Atoi(id); err != nil {
		return false
	} else if result := r.db.Find(&module,i); result.Error != nil {
		return false
	} else {
		return true
	}
}

func (r *MySqlRepository) CheckIfModuleVisitExists(id string) bool {
	var visit persistence.ModuleVisit
	if result := r.db.Where("module = ?", id).First(&visit); result.Error != nil {
		return false
	} else {
		return true
	}
}
