package persistence

//User type for User Handlers
type User struct {
	ID uint `gorm:"number(3);PRIMARY_KEY;AUTO_INCREMENT"`
	Name string
	Email string
	Password string `json:"Password"`
}

