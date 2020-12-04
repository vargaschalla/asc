package models

import (
	"gorm.io/gorm"
)

type Sesiones struct {
	gorm.Model
	Nombre      string `json:"nombre"`
	Titulo      string `json:"titulo"`
	Descripcion string `json:"descripcion"`
	Estado      string `json:"estado"`
}
