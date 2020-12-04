package models
import (
	"gorm.io/gorm"
)
type Tareas struct {
    gorm.Model
	Curso 		string `json:"curso"`
	Titulo	 	string `json:"titulo"`
	Nota		string `json:"nota"`
	Estado		string `json:"estado"`
}