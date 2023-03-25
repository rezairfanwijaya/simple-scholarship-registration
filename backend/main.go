package main

import (
	"log"
	"net/http"
	"serkom/admin"
	"serkom/approval"
	"serkom/auth"
	"serkom/database"
	"serkom/handler"
	"serkom/helper"
	"serkom/mahasiswa"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

	// auth
	authService := auth.NewAuthService()

	// approval
	repoApproval := approval.NewRepository(conn)
	serviceApproval := approval.NewService(repoApproval)

	// mahasiswa
	repoMahasiswa := mahasiswa.NewRepository(conn)
	serviceMahasiswa := mahasiswa.NewService(repoMahasiswa, serviceApproval)
	hanlderMahasiswa := handler.NewHandlerMahasiswa(serviceMahasiswa)

	// admin
	// admin
	repoAdmin := admin.NewRepository(conn)
	serviceAdmin := admin.NewService(repoAdmin, serviceApproval)
	handlerAdmin := handler.NewHandlerAdmin(serviceAdmin, authService)

	// endpoint
	router.GET("/mahasiswa/:email", hanlderMahasiswa.GetByEmail)
	router.POST("/mahasiswa/daftar", hanlderMahasiswa.DaftarBeasiswa)
	router.GET("/mahasiswa/approval", hanlderMahasiswa.GetAllApprovals)
	router.POST("/admin/login", handlerAdmin.Login)
	router.GET("/admin/status/show", handlerAdmin.GetTotalStatus)
	router.PUT("/admin/status/:id", handlerAdmin.ChangeStatusApporval)
	router.Static("/image", "./images")


	// start server
	if err := router.Run(":8989"); err != nil {
		log.Fatal(err.Error())
	}
}

// function middleware auth
func authMiddleware(authService auth.Service, adminService admin.Service, typeUser string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil nilai header yang sudah kita set namanya Authorization
		authHeader := c.GetHeader("Authorization")

		// cek apakah nilai authorization memiliki Bearer
		// karena nanti kita akan set nilai token seperti ini "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		if !strings.Contains(authHeader, "Bearer") {
			respons := helper.GenerateResponse(
				"gagal",
				http.StatusUnauthorized,
				"Unauthorized",
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// lalu kita pisahkan tokennya berdasarkan spasi
		// before "Bearer djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"
		// after ["Bearer"] ["djfkfbnfkjbnfkjgbnfkyreguryhvbfdhvbfbvdhbf"]
		tokenString := ""
		arraytoken := strings.Split(authHeader, " ")
		if len(arraytoken) == 2 {
			tokenString = arraytoken[1]
		}

		// validasi token
		token, err := authService.VerifyToken(tokenString)
		if err != nil {
			respons := helper.GenerateResponse(
				"gagal",
				http.StatusUnauthorized,
				"Unauthorized",
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// ambil data dalam token
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respons := helper.GenerateResponse(
				"gagal",
				http.StatusUnauthorized,
				"Unauthorized",
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
			return
		}

		// data id user diambil
		userId := int(claim["user_id"].(float64))

		if typeUser == "admin" {
			// data user diambil berdasarkan id
			user, err := adminService.GetUserByID(userId)
			if err != nil {
				respons := helper.GenerateResponse(
					"gagal",
					http.StatusUnauthorized,
					"Unauthorized",
				)
				c.AbortWithStatusJSON(http.StatusUnauthorized, respons)
				return
			}

			// set context
			c.Set("currentAdmin", user)
		}
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
