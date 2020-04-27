// Student Models, with JSON names and validation tags
package persistence

type Student struct {
	ID uint `json:"id,omitempty,string" validate:"required"`
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=320"`
	Semester string `json:"semester,omitempty" validate:"required"`
	Password string `json:"-"`
	Degree string `json:"degree,omitempty" validate:"required,numeric"`
}
// Model without ID for Creation of Module Visit
type CreateStudentBody struct {
	Email string `json:"email,omitempty" validate:"required,email,min=6,max=320"`
	Password string `json:"password,omitempty" validate:"required,min=10"`
	Semester string `json:"semester,omitempty" validate:"required"`
	Degree string `json:"degree,omitempty" validate:"required,numeric"`
}


