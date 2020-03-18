package persistence

type Repository interface {
	GetAllStudents() []Student
	FindOne(email, password string) map[string]interface{}
	CreateStudent(user *Student) (*Student,error)
	GetStudentById(id string) Student
	UpdateStudent(id string, student *Student) *Student
	SaveAllModules([]Module)
	GetAllModules() []Module
	GetModuleById(id string) Module
}

type DefaultRepository struct {
	Users []Student
	Modules []Module
}
