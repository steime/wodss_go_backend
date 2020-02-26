package persistence

type Repository interface {
	GetAllUsers() ([]User, error)

}

type DefaultRepository struct {
	Users []User
}
