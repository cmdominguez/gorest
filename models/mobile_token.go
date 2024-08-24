package models

import (
	"gorm.io/gorm"
)

type MobileToken struct {
	gorm.Model
	Token string
}
