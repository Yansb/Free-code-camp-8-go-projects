package models

import "gorm.io/gorm"

type Data struct {
	ID          uint `gorm:"primaryKey"`
	Idade       string
	Trabalho    string
	Numero      string
	Graduacao   string
	OutroNumero string
	EstadoCivil string
	gorm.Model
}
