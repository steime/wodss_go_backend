package mySQL

import (
	"github.com/steime/wodss_go_backend/persistence"
	"strconv"
)

func (r *MySqlRepository) SaveAllDegrees(degrees []persistence.Degree){
	for _ , d := range degrees {
		r.db.Create(&d)
	}
}

func (r *MySqlRepository) UpdateAllDegrees(degrees []persistence.Degree){
	for _ , d := range degrees {
		r.db.Save(&d)
	}
}

func (r *MySqlRepository) GetAllDegrees() ([]persistence.Degree,error) {
	var degrees []persistence.Degree
	if result := r.db.Preload("Groups").Preload("ProfilesByDegree").Find(&degrees); result.Error != nil {
		return degrees,result.Error
	} else {
		return degrees,nil
	}
}

func (r *MySqlRepository) GetDegreeById(id string) (persistence.Degree,error) {
	var degree persistence.Degree
	if i,err := strconv.Atoi(id); err != nil {
		return degree, err
	} else if result := r.db.Preload("Groups").Preload("ProfilesByDegree").Find(&degree,i); result.Error != nil {
		return degree,result.Error
	} else {
		return degree,nil
	}

}
