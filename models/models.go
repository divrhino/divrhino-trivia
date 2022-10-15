package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
