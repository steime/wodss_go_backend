package persistence

type Repository interface {
	GetAllUsers() []User
	FindOne(email, password string) map[string]interface{}
	CreateUser(user *User)
}

type DefaultRepository struct {
	Users []User
}
