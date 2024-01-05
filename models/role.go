package models

type Role struct {
	ID    uint    `gorm:"primary_key" json:"id"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_roles;" json:"users"`
}
