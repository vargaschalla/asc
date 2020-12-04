package models

import "gorm.io/gorm"

type Persona struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Paterno  string `json:"paterno"`
	Materno  string `json:"materno"`
	Usuario  string `json:"usuario"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Edad     int    `json:"edad"`
	DNI      int    `json:"dni"`
}
