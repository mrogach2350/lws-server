package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	Label  string  `json:"label"`
	Value  string  `json:"value"`
	Movies []Movie `json:"movies"`
}
