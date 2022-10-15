package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string
	Answer   string
}
