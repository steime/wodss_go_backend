package mySQL

import (
	"fmt"
	"github.com/steime/wodss_go_backend/persistence"
)

func (r *MySqlRepository) SaveAllDegrees(degrees []persistence.Degree){
	for _ , d := range degrees {
		fmt.Println(d)
		r.db.Create(&d)
	}
}
