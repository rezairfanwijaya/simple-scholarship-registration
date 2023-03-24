package approval

type Approval struct {
	ID     int `json:"id" gorm:"primaryKey"`
	UserID int `json:"user_id"`
	Status int `json:"status"`
}
