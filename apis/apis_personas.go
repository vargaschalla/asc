package apis

import (
	"net/http"
	"pelao/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PersonasLista(c *gin.Context) {
	var lis []models.Persona

	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}

func PersonasCreate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	var d models.Persona
	//d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}

func PersonasUpdate(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Persona
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Nombre = c.PostForm("nombre")
	d.Paterno = c.PostForm("paterno")
	d.Materno = c.PostForm("materno")
	d.Usuario = c.PostForm("usuario")
	d.Email = c.PostForm("email")
	d.Password = c.PostForm("password")
	d.Edad = 0
	d.DNI = 0
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}
func PersonasDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Persona

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
