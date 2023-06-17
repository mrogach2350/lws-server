package models

import "time"

type Genre struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int    `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
}
