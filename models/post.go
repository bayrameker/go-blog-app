package models

import "time"

type Post struct {
	ID          uint        `gorm:"primary_key" json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Content     string      `json:"content"`
	AuthorID    uint        `json:"author_id"` // Foreign key (belongs to), tag `json:"author_id"` is optional
	Categories  []*Category `gorm:"many2many:post_categories;"`
	Comments    []Comment   `json:"comments"`
	CreatedAt   *time.Time  `json:"created_at"`
	UpdatedAt   *time.Time  `json:"updated_at"`
	DeletedAt   *time.Time  `json:"deleted_at"`
}
