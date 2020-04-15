package persistence

type Repository interface {
	GetAllStudents() []Student
	FindOne(email, password string) map[string]interface{}
	CreateStudent(user *Student) (*Student,error)
	GetStudentById(id string) (Student,error)
	DeleteStudent(id string) error
	UpdateStudent(id string, student *Student) (*Student,error)
	GetAllModules() ([]Module,error)
	SaveAllModules([]Module)
	GetModuleById(id string) (Module,error)
	GetAllModuleGroups() ([]ModuleGroup,error)
	SaveAllModuleGroups([]ModuleGroup)
	GetModuleGroupById(id string) (ModuleGroup,error)
	GetAllDegrees() ([]Degree,error)
	SaveAllDegrees([]Degree)
	GetDegreeById(id string) (Degree,error)
	CreateModuleVisit(visit *ModuleVisit) (*ModuleVisit,error)
	GetAllModuleVisits(studentId string) ([]ModuleVisit,error)
	GetModuleVisitById(moduleId string,studentId string) (ModuleVisit,error)
	UpdateModuleVisit(visit *ModuleVisit) (*ModuleVisit,error)
	DeleteModuleVisit(visitId string, studentId string) error
	ForgotPassword(mail string) error
	ResetPassword(mail string, password string) error
}

type DefaultRepository struct {
	Users []Student
	Modules []Module
}
