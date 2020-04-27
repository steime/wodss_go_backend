// Auth Models, with JSON names and validation tags
package persistence

type PasswordResetBody struct {
	ForgotToken string 	`json:"forgotToken" validate:"required"`
	Email string 		`json:"email" validate:"required,email,min=6,max=320"`
	Password string 	`json:"password" validate:"required,min=10"`
}

type LoginBody struct {
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type ForgotTokenContext struct {
	ForgotToken string
}

type TokenPair struct {
	Token string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

type TokenRequestBody struct {
	RefreshToken string `json:"refreshToken"`
}
