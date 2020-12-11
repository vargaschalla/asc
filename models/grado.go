package models

import "gorm.io/gorm"

type Grado struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}
