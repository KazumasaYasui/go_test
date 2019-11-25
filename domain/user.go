package domain

type User struct {
	ID    int
	Name  string `json:"name" gorm:"size:255"`
	Email string `json:"email"`
}

type Users []User
