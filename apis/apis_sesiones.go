package apis

import (
	"net/http"
	"pelao/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SesionesGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var ses models.Sesiones
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&ses, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &ses)
}

func SesionesIndex(c *gin.Context) {
	var lis []models.Sesiones
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}

func SesionesPost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var ses models.Sesiones
	//tar := models.Sesiones{Curso: c.PostForm("curso"), Titulo: c.PostForm("titulo"), Nota: c.PostForm("nota"), Estado: c.PostForm("estado")}
	if err := c.BindJSON(&ses); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&ses)
	c.JSON(http.StatusOK, &ses)
}

func SesionesPut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var ses models.Sesiones
	if err := conn.First(&ses, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ses.Nombre = c.PostForm("nombre")
	ses.Titulo = c.PostForm("titulo")
	ses.Descripcion = c.PostForm("descripcion")
	ses.Estado = c.PostForm("estado")
	conn.Save(&ses)
	c.JSON(http.StatusOK, &ses)
}

func SesionesDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var ses models.Sesiones
	if err := conn.Where("id = ?", id).First(&ses).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&ses)
}
