package persistence

type Repository interface {
	GetAllUsers() []User
	AddUser(user *User)
}

type DefaultRepository struct {
	Users []User
}
