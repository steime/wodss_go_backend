package persistence

import (
	"encoding/json"
)

type User struct {
	Id int 'json:"Id"'
	Name string 'json:"Name"'
}

var Users []User