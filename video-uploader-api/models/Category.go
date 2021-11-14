package models

type Category struct {
	Base
	Name string `gorm:"not null;" json:"name"`
}
