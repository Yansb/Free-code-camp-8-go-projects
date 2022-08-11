package models

import "gorm.io/gorm"

type Data struct {
	ID       uint `gorm:"primaryKey"`
	WO       string
	District string
	LeadTech string
	Service  string
	Techs    string
	LbrHrs   string
	PartsCst string
	Payment  string
	gorm.Model
}
