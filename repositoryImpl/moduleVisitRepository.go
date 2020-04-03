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
	} else if r.CheckIfModuleVisitExists(studentID,moduleID) {
		return visit,errors.New("moduleVisit for this module exists")
	} else {
		if result := r.db.Create(&visit); result.Error != nil {
			return visit,result.Error
		} else {
			return visit,nil
		}
	}
}

func (r *MySqlRepository) GetAllModuleVisits(studentId string) ([]persistence.ModuleVisit,error) {
	var visits []persistence.ModuleVisit
	if result := r.db.Where("student = ?",studentId).Find(&visits); result.Error != nil {
		return visits,result.Error
	} else {
		return visits,nil
	}
}

func (r *MySqlRepository) GetModuleVisitById(visitId string, studentId string) (persistence.ModuleVisit,error) {
	var visit persistence.ModuleVisit
	if result := r.db.Where("id = ? AND student = ?", visitId,studentId).First(&visit); result.Error != nil {
		return visit,result.Error
	} else {
		return visit,nil
	}
}

func (r *MySqlRepository) UpdateModuleVisit(visit *persistence.ModuleVisit) (*persistence.ModuleVisit,error) {
	if !r.CheckIfModuleExists(visit.Module) {
		return visit,errors.New("module not existing")
	} else if result := r.db.Model(&visit).Omit("id").Update(persistence.ModuleVisit{
		Grade:     visit.Grade,
		State:     visit.State,
		Student:   visit.Student,
		Module:    visit.Module,
		Semester:  visit.Semester,
		Weekday:   visit.Weekday,
		TimeStart: visit.TimeStart,
		TimeEnd:   visit.TimeEnd,
	}); result.Error !=nil {
		return visit,result.Error
	} else {
		return visit,nil
	}
}

func (r *MySqlRepository) DeleteModuleVisit(visitId string, studentId string) error {
	var visit persistence.ModuleVisit
	if result := r.db.Where("id = ? AND student = ?",visitId,studentId).First(&visit); result.Error != nil {
		return result.Error
	} else if result = r.db.Delete(visit); result.Error != nil {
		return result.Error
	} else {
		return nil
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

func (r *MySqlRepository) CheckIfModuleVisitExists(studentID uint,moduleID string) bool {
	var visit persistence.ModuleVisit
	if result := r.db.Where("module = ? AND student = ?", moduleID,studentID).First(&visit); result.Error != nil {
		return false
	} else {
		return true
	}
}
