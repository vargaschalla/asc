package apis

import (
	"net/http"
	"pelao/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RolLista(c *gin.Context) {
	var lis []models.Rol

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista",
		"r":   lis,
	})
}

func RolCreate(c *gin.Context) {
	var p models.Rol
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	if err := c.BindJSON(&p); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	conn.Create(&p)
	c.JSON(http.StatusOK, &p)
}

func RolUpdate(c *gin.Context) {
	var d models.Rol
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Nombre = c.PostForm("nombre")
	d.Codigo = c.PostForm("codigo")
	d.Estado = 0
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}
func RolDelete(c *gin.Context) {
	db, _ := c.Get("db")

	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Rol

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
