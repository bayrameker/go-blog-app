package models

type Comment struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Content string `json:"content"`
	PostID  uint   `json:"post_id"`
	UserID  uint   `json:"user_id"`
}
