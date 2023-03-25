package approval

type InputNewApproval struct {
	UserID        int    `json:"user_id"`
	Status        int    `json:"status"`
	JenisBeasiswa string `json:"jenis_beasiswa"`
}
