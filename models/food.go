package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Name            string `json:"name"`
	CategoryName    string
	Category        Category `gorm:"references:Name"`
	SubCategoryName string
	SubCategory     Category       `gorm:"references:Name"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted"`
}
