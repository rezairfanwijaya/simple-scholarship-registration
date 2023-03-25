package approval

import "gorm.io/gorm"

type IRespoistory interface {
	Save(approval Approval) (Approval, error)
	FindAll() ([]Approval, error)
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
