package persistence

type Repository interface {
	GetAllUsers() []User
	FindOne(email, password string) map[string]interface{}
	CreateUser(user *User) (*User,error)
	FindById(id string) User
	//GetAllModules()
	SaveAllModules([]Module)
	GetAllModules() []Module
	GetModuleById(id string) Module
}

type DefaultRepository struct {
	Users []User
	Modules []Module
}
