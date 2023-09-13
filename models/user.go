package models

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Password []byte `json:"password"`
	Email    string `gorm:"unique" json:"email"`
}