package admin

type ResponseAdminSuccessLogin struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Role  string `json:"role"`
}

type responseTotalStatus struct {
	Pending int64 `json:"pending"`
	Approve int64 `json:"approve"`
}

func FormatResponseAdminSuccessLogin(admin Admin, token string) *ResponseAdminSuccessLogin {
	return &ResponseAdminSuccessLogin{
		ID:    admin.ID,
		Email: admin.Email,
		Role:  admin.Role,
		Token: token,
	}
}

func FormatTotalStatus(totalPending, totalAppove int64) *responseTotalStatus {
	return &responseTotalStatus{
		Pending: totalPending,
		Approve: totalAppove,
	}
}
