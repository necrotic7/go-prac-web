package models

type User struct{
	ID uint `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReigisterUser struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	ConfirmPwd string `json:"confirmpwd" binding:"required"`
}