package models

type User struct {
	ID       uint      `gorm:"primaryKey"`
	Login    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Posts    []Post    `gorm:"foreignKey:UserID"`
	Comments []Comment `gorm:"foreignKey:UserID"`
}
