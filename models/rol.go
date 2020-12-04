package models

import "gorm.io/gorm"

type Rol struct {
	gorm.Model
	Nombre string `json:"nombre"`
	Codigo string `json:"codigo"`
	Estado int    `json:"estado"`
}
