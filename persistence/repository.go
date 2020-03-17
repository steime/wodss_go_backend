package persistence

type Repository interface {
	GetAllStudents() []Student
	FindOne(email, password string) map[string]interface{}
	CreateStudent(user *Student) (*Student,error)
	FindById(id string) Student
	//GetAllModules()
	SaveAllModules([]Module)
	GetAllModules() []Module
	GetModuleById(id string) Module
}

type DefaultRepository struct {
	Users []Student
	Modules []Module
}
