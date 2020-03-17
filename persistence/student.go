package persistence

//User type for User Handlers
type Student struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Semester string `json:"semester"`
	Password string `json:"-"`
	Degree string `json:"degree"`
}


