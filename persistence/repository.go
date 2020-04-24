package persistence

type Repository interface {
	FindOne(email, password string) (TokenPair,error)
	CreateStudent(user *Student) (*Student,error)
	GetStudentById(id string) (Student,error)
	DeleteStudent(id string) error
	UpdateStudent(id string, student *Student) (*Student,error)
	GetAllModules() ([]Module,error)
	SaveAllModules([]Module)
	UpdateAllModules([]Module)
	GetModuleById(id string) (Module,error)
	GetAllModuleGroups() ([]ModuleGroup,error)
	SaveAllModuleGroups([]ModuleGroup)
	UpdateAllModuleGroups([]ModuleGroup)
	GetModuleGroupById(id string) (ModuleGroup,error)
	GetAllDegrees() ([]Degree,error)
	SaveAllDegrees([]Degree)
	UpdateAllDegrees([]Degree)
	GetDegreeById(id string) (Degree,error)
	CreateModuleVisit(visit *ModuleVisit) (*ModuleVisit,error)
	GetAllModuleVisits(studentId string) ([]ModuleVisit,error)
	GetModuleVisitById(moduleId string,studentId string) (ModuleVisit,error)
	UpdateModuleVisit(visit *ModuleVisit) (*ModuleVisit,error)
	DeleteModuleVisit(visitId string, studentId string) error
	ForgotPassword(mail string) error
	ResetPassword(mail string, password string) error
	GetAllProfiles() ([]Profile,error)
	SaveAllProfiles([]Profile)
	UpdateAllProfiles([]Profile)
	GetProfileById(id string) (Profile,error)
}
