package handler

import (
	"serkom/helper"
	"serkom/mahasiswa"

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
