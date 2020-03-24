package persistence

type Repository interface {
	GetAllStudents() []Student
	FindOne(email, password string) map[string]interface{}
	CreateStudent(user *Student) (*Student,error)
	GetStudentById(id string) (Student,error)
	DeleteStudent(id string) error
	UpdateStudent(id string, student *Student) (*Student,error)
	GetAllModules() []Module
	SaveAllModules([]Module)
	GetModuleById(id string) Module
	GetAllModuleGroups() ([]ModuleGroup,error)
	SaveAllModuleGroups([]ModuleGroup)
	GetModuleGroupById(id string) (ModuleGroup,error)
}

type DefaultRepository struct {
	Users []Student
	Modules []Module
}
