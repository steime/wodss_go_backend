package persistence

type Student struct {
	ID uint `json:"id,omitempty,string" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email"`
	Semester string `json:"semester,omitempty" validate:"required"`
	Password string `json:"-"`
	Degree string `json:"degree,omitempty" validate:"required"`
}


