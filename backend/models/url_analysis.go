package models

import "gorm.io/gorm"

type URLAnalysis struct {
	gorm.Model
	URL          string `gorm:"not null"`
	HTMLVersion  string
	Title        string
	Headings     string
	Links        string
	BrokenLinks  string
	HasLoginForm bool
}
