package admin

type InputLoginAdmin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type InputChangeStatusApproval struct {
	Status int `json:"status"`
}
