package Model

import "github.com/jinzhu/gorm"

type Example struct {
	gorm.Model
	
	Text string `json:"text"`
}
