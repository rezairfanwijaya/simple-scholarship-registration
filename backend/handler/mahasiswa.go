package handler

import (
	"fmt"
	"log"
	"net/http"
	"serkom/helper"
	"serkom/mahasiswa"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handlerMahasiswa struct {
	mahasiswaService mahasiswa.IService
}

func NewHandlerMahasiswa(mahasiswaService mahasiswa.IService) *handlerMahasiswa {
	return &handlerMahasiswa{mahasiswaService}
}

func (h *handlerMahasiswa) GetByEmail(c *gin.Context) {
	// get email from param
	email := c.Param("email")

	// call service
	mahaisiswaByEmail, code, err := h.mahasiswaService.GetByEmail(email)
	if err != nil {
		response := helper.GenerateResponse(
			"error",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	// format mahasiswa
	mahasiswaForamtted := mahasiswa.FormatterMahasiswa(mahaisiswaByEmail)

	response := helper.GenerateResponse(
		"sukses",
		code,
		mahasiswaForamtted,
	)

	c.JSON(code, response)
}

func (h *handlerMahasiswa) DaftarBeasiswa(c *gin.Context) {
	// get data from form
	nama := c.PostForm("nama")
	email := c.PostForm("email")
	telepon := c.PostForm("telepon")
	semester := c.PostForm("semester")
	beasiswa := c.PostForm("beasiswa")
	ipk := c.PostForm("ipk")

	berkas, err := c.FormFile("berkas")
	if err != nil {
		response := helper.GenerateResponse(
			"error",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get path for berkas
	path, fileName, code, err := helper.ServeBerkas(berkas.Filename)
	if err != nil {
		response := helper.GenerateResponse(
			"error",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	// save berkas to local
	if err := c.SaveUploadedFile(berkas, path); err != nil {
		response := helper.GenerateResponse(
			"failed",
			http.StatusBadRequest,
			err.Error(),
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// generate url image for client
	berkasURL := fmt.Sprintf("http://localhost:6776/image/%v", fileName)

	// set input for service
	semesterInNumber, _ := strconv.Atoi(semester)
	ipkInFloat, _ := strconv.ParseFloat(ipk, 0)
	input := mahasiswa.InputDaftarBeasiswa{
		Nama:          nama,
		Email:         email,
		Telepon:       telepon,
		Semester:      semesterInNumber,
		IPK:           ipkInFloat,
		BerkasURL:     berkasURL,
		JenisBeassiwa: beasiswa,
	}

	log.Println(input)
}
