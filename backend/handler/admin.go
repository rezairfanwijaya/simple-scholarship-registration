package handler

import (
	"net/http"
	"serkom/admin"
	"serkom/auth"
	"serkom/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HandlerAdmin struct {
	adminService admin.IService
	authService  auth.IService
}

func NewHandlerAdmin(adminService admin.IService, authService auth.IService) *HandlerAdmin {
	return &HandlerAdmin{adminService, authService}
}

func (h *HandlerAdmin) Login(c *gin.Context) {
	var input admin.InputLoginAdmin

	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	adminLoggedin, err := h.adminService.Login(input)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// generate jwt
	token, err := h.authService.GenerateToken(adminLoggedin.ID)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	adminFormatted := admin.FormatResponseAdminSuccessLogin(adminLoggedin, token)

	response := helper.GenerateResponse(
		"sukses",
		http.StatusOK,
		adminFormatted,
	)

	c.JSON(http.StatusOK, response)
}

func (h *HandlerAdmin) GetTotalStatus(c *gin.Context) {
	totalPending, totalApprove := h.adminService.GetTotalStatus()

	statusFotmatted := admin.FormatTotalStatus(totalPending, totalApprove)

	response := helper.GenerateResponse(
		"sukses",
		http.StatusOK,
		statusFotmatted,
	)

	c.JSON(http.StatusOK, response)
}

func (h *HandlerAdmin) ChangeStatusApporval(c *gin.Context) {
	// get id from param
	IDString := c.Param("id")
	IDNumber, err := strconv.Atoi(IDString)
	if err != nil || IDNumber == 0 {
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			"ID harus integer dan lebih besar dari 0",
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// binding
	var input admin.InputChangeStatusApproval
	if err := c.BindJSON(&input); err != nil {
		errBinding := helper.GenerateErrorBinding(err)
		response := helper.GenerateResponse(
			"gagal",
			http.StatusBadRequest,
			errBinding,
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// call service
	_, code, err := h.adminService.ChangeStatus(IDNumber, input.Status)
	if err != nil {
		response := helper.GenerateResponse(
			"gagal",
			code,
			err.Error(),
		)

		c.JSON(code, response)
		return
	}

	response := helper.GenerateResponse(
		"sukses",
		code,
		"sukses update status",
	)

	c.JSON(code, response)
}
