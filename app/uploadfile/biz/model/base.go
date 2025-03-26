package model

import "gorm.io/gorm"

type Base struct {
	gorm.Model
  	ID        uint           `gorm:"primaryKey"`
}