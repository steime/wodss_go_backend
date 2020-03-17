package persistence

//User type for User Handlers
type Student struct {
	ID uint `json:"id"`
	Semester string `json:"semester"`
	Email string `json:"email"`
	Password string `json:"-"`
	Degree string `json:"degree"`
}


