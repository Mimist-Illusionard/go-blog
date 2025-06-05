package models

type Comment struct {
	ID    uint `gorm:"primaryKey"`
	Owner string
	Text  string
}
