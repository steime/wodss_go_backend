package persistence

type Repository interface {
	GetAllUsers() []User
	FindOne(email, password string) map[string]interface{}
	CreateUser(user *User) (*User,error)
	FindById(id string) User
}

type DefaultRepository struct {
	Users []User
}
