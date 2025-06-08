package models

type Post struct {
	ID       uint      `gorm:"primaryKey"`
	Text     string    `gorm:"not null"`
	UserID   uint      `gorm:"index"`
	User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}
