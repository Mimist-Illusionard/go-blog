package models

type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	Text   string `gorm:"not null"`
	UserID uint   `gorm:"index"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PostID uint   `gorm:"index"`
	Post   Post   `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
