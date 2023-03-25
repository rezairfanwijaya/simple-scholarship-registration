package main

import (
	"log"
	"serkom/approval"
	"serkom/database"
	"serkom/handler"
	"serkom/helper"
	"serkom/mahasiswa"

	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// init router
	router := gin.Default()

	// allow cors
	router.Use(CorsMiddleware())

	// approval
	repoApproval := approval.NewRepository(conn)
	serviceApproval := approval.NewService(repoApproval)

	// mahasiswa
	repoMahasiswa := mahasiswa.NewRepository(conn)
	serviceMahasiswa := mahasiswa.NewService(repoMahasiswa, serviceApproval)
	hanlderMahasiswa := handler.NewHandlerMahasiswa(serviceMahasiswa)

	// endpoint
	router.GET("/mahasiswa/:email", hanlderMahasiswa.GetByEmail)
	router.POST("/mahasiswa/daftar", hanlderMahasiswa.DaftarBeasiswa)
	router.GET("/mahasiswa/approval", hanlderMahasiswa.GetAllApprovals)
	router.Static("/image", "./images")

	env, err := helper.GetENV(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	// start server
	if err := router.Run(env["DOMAIN"]); err != nil {
		log.Fatal(err.Error())
	}
}

// CORS
func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}
