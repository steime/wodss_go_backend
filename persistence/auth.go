package persistence

type PasswordResetBody struct {
	ForgotToken string 	`json:"forgotToken" validate:"required"`
	Email string 		`json:"email" validate:"required,email,min=6,max=320"`
	Password string 	`json:"password" validate:"required,min=10"`
}
