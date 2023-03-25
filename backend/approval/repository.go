package approval

import "gorm.io/gorm"

type IRespoistory interface {
	Save(approval Approval) (Approval, error)
	FindAll() ([]Approval, error)
	CountStatus(pending, approve int) (totalPending int64, totalApprove int64)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(approval Approval) (Approval, error) {
	if err := r.db.Create(&approval).Error; err != nil {
		return approval, err
	}

	return approval, nil
}

func (r *repository) FindAll() ([]Approval, error) {
	var approvals []Approval
	if err := r.db.Order("id desc").Find(&approvals).Error; err != nil {
		return approvals, err
	}

	return approvals, nil
}

func (r *repository) CountStatus(pending, approve int) (totalPending int64, totalApprove int64) {
	r.db.Model(&Approval{}).Where("status = ? ", pending).Count(&totalPending)
	r.db.Model(&Approval{}).Where("status = ? ", approve).Count(&totalApprove)

	return totalPending, totalApprove
}
