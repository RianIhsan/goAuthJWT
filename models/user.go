package models

type User struct {
	Id       uint   `json:"id" gorm:"primarykey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
