package models

import "time"

type User struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string
	Posts     []Post     `gorm:"foreignKey:AuthorID" json:"posts"`
	Comments  []Comment  `json:"comments"`
	Roles     []*Role    `gorm:"many2many:user_roles;" json:"roles"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
