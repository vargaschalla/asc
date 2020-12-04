package routers

import (
	"pelao/apis"
	"pelao/models"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	ses := r.Group("/ses")
	{
		ses.GET("/sesiones/:id", apis.SesionesGetId)
		ses.GET("/sesiones", apis.SesionesIndex)
		ses.POST("/sesiones", apis.SesionesPost)
		ses.PUT("/sesiones/:id", apis.SesionesPut)
		ses.DELETE("/sesiones/:id", apis.SesionesDelete)
	}

	tar := r.Group("/tar")
	{
		tar.GET("/tareas/:id", apis.TareasGetId)
		tar.GET("/tareas", apis.TareasIndex)
		tar.POST("/tareas", apis.TareasPost)
		tar.PUT("/tareas/:id", apis.TareasPut)
		tar.DELETE("/tareas/:id", apis.TareasDelete)
	}

	us := r.Group("/us")
	{
		us.GET("/user", apis.UsersIndex)
		us.POST("/user", apis.UsersCreate)
		us.GET("/user/:id", apis.UsersGet)
		us.PUT("/user/:id", apis.UsersUpdate)
		us.DELETE("/user/:id", apis.UsersDelete)
		us.POST("/login", apis.UsersLogin)
		us.POST("/logout", apis.UsersLogout)
	}
	per := r.Group("/per")
	{
		per.GET("/persons", apis.PersonasLista)
		per.POST("/persons", apis.PersonasCreate)
		per.PUT("/persons/:id", apis.PersonasUpdate)
		per.DELETE("/persons/:id", apis.PersonasDelete)

	}
	rol := r.Group("/r")
	{
		rol.GET("/rol", apis.RolLista)
		rol.POST("/rol", apis.RolCreate)
		rol.PUT("/rol/:id", apis.RolUpdate)
		rol.DELETE("/rol/:id", apis.RolDelete)

	}

	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "root:ANTHONY2020ma@tcp(localhost:3306)/pruebas?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	conn.AutoMigrate(&models.Tareas{}, &models.Sesiones{}, &models.User{}, &models.Persona{}, &models.Rol{})

	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin", "http://localhost","http://localhost:8080")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		//fmt.Printf("Bearer (%v) \n", token)
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
