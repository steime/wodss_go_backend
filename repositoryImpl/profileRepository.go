package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
	"strconv"
)

func (r *MySqlRepository) SaveAllProfiles(degrees []persistence.Profile) {
	for _, d := range degrees {
		fmt.Println(d)
		r.db.Create(&d)
	}
}

func (r *MySqlRepository) GetAllProfiles() ([]persistence.Profile,error) {
	var profiles []persistence.Profile
	if result := r.db.Preload("ListOfModules").Find(&profiles); result.Error != nil {
		return profiles,result.Error
	} else {
		return profiles,nil
	}
}

func (r *MySqlRepository) GetProfileById(id string) (persistence.Profile,error) {
	var profile persistence.Profile
	if i,err := strconv.Atoi(id); err != nil {
		return profile, err
	} else if result := r.db.Preload("ListOfModules").Find(&profile,i); result.Error != nil {
		return profile,result.Error
	} else {
		return profile,nil
	}

}
