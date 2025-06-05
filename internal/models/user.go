package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Login    string
	Password string
}
