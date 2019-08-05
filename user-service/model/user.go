package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Company  string `json:"company"`
	Email    string `json:"email" gorm:"type:text;UNIQUE"`
	Password string `json:"password"`
}

func (u *User) IsValid() bool {
	if u.Email == "" || u.Password == "" {
		return false
	}
	return true
}
