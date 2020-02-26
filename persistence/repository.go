package persistence

type Repository interface {
	//GetAllUsers() ([]User, error)
	AddUser(user *User)
}

type DefaultRepository struct {
	Users []User
}
