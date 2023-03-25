package admin

import "gorm.io/gorm"

type IRepository interface {
	FindByEmail(email string) (Admin, error)
	FindByID(userID int) (Admin, error)
}

type Respository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Respository {
	return &Respository{db}
}

func (r *Respository) FindByEmail(email string) (Admin, error) {
	var admin Admin
	if err := r.db.Where("email = ?", email).Find(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}

func (r *Respository) FindByID(userID int) (Admin, error) {
	var admin Admin
	if err := r.db.Where("id = ?", userID).Find(&admin).Error; err != nil {
		return admin, err
	}

	return admin, nil
}
