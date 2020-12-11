package models

import "gorm.io/gorm"

type Nivel struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}
