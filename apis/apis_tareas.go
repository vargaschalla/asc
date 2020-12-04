package apis

import (
	"net/http"
	"pelao/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TareasGetId(c *gin.Context) {
	id := c.Params.ByName("id")
	var tar models.Tareas
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	if err := conn.First(&tar, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &tar)
}

func TareasIndex(c *gin.Context) {
	var lis []models.Tareas
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	c.JSON(http.StatusOK, lis)
}

func TareasPost(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var tar models.Tareas
	//tar := models.Tareas{Curso: c.PostForm("curso"), Titulo: c.PostForm("titulo"), Nota: c.PostForm("nota"), Estado: c.PostForm("estado")}
	if err := c.BindJSON(&tar); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&tar)
	c.JSON(http.StatusOK, &tar)
}

func TareasPut(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var tar models.Tareas
	if err := conn.First(&tar, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tar.Curso = c.PostForm("curso")
	tar.Titulo = c.PostForm("titulo")
	tar.Nota = c.PostForm("nota")
	tar.Estado = c.PostForm("estado")
	conn.Save(&tar)
	c.JSON(http.StatusOK, &tar)
}

func TareasDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var tar models.Tareas
	if err := conn.Where("id = ?", id).First(&tar).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&tar)
}
