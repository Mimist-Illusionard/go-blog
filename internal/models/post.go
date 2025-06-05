package models

type Post struct {
	ID   uint `gorm:"primaryKey"`
	Text string
}
