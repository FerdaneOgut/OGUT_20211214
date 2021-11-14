package models

type Video struct {
	Base
	UID          string `gorm:"not null;" json:"uid"`
	Title        string `gorm:"not null;" json:"title"`
	CategoryID   int    `gorm:not null; json "categoryID"`
	Thumbnail64  []byte `gorm:"not null;" json:"thumbnail64"`
	Thumbnail128 []byte `gorm:"not null;" json:"thumbnail128"`
	Thumbnail256 []byte `gorm:"not null;" json:"thumbnail256"`
	Path         string `gorm:"not null;" json:"path"`
	Category     Category
}
