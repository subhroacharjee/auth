package user

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
	IsVerified bool   `json:"isVerified"`
}

func (u User) ToJson() map[string]string {
	return map[string]string{
		"id":          fmt.Sprintf("%d", u.ID),
		"name":        u.Name,
		"email":       u.Email,
		"is_verified": fmt.Sprintf("%t", u.IsVerified),
	}
}
