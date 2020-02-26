package persistence

import "github.com/jinzhu/gorm"

//User type for User Handlers
type User struct {
	gorm.Model
	Name string
}

