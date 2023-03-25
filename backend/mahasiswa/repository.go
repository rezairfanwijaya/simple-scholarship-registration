package mahasiswa

import "gorm.io/gorm"

type IRespoistory interface {
	FindByEmail(email string) (Mahasiswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByEmail(email string) (Mahasiswa, error) {
	var mahasiswa Mahasiswa
	if err := r.db.Where("email = ?", email).Find(&mahasiswa).Error; err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}
