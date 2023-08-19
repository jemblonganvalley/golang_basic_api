package api

import "gorm.io/gorm"

type UserType struct {
	Email string 
	Username string 
	Password *string 
}

type UserResponse struct {
	ID int
	Email string `json:"email"`
	Username string `json:"username"`
}

type UserModel struct {
	gorm.Model
	Email string
	Password string
	Username string
}