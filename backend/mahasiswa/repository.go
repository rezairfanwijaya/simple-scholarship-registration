package mahasiswa

import "gorm.io/gorm"

type IRespoistory interface {
	FindByID(ID int) (Mahasiswa,error)
	FindByEmail(email string) (Mahasiswa, error)
	Update(mahasiswa Mahasiswa) (Mahasiswa, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(ID int) (Mahasiswa,error) {
	var mahasiswa Mahasiswa
	if err := r.db.Where("id = ?", ID).Find(&mahasiswa).Error; err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}

func (r *repository) FindByEmail(email string) (Mahasiswa, error) {
	var mahasiswa Mahasiswa
	if err := r.db.Where("email = ?", email).Find(&mahasiswa).Error; err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}

func (r *repository) Update(mahasiswa Mahasiswa) (Mahasiswa, error) {
	if err := r.db.Save(&mahasiswa).Error; err != nil {
		return mahasiswa, err
	}

	return mahasiswa, nil
}
