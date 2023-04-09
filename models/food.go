package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	Name            string         `json:"name"`
	CategoryName    string         `json:"categoryName"`
	SubCategoryName string         `json:"subcategoryName"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted"`
}
