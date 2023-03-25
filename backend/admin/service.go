package admin

import (
	"fmt"
	"serkom/approval"
	"serkom/helper"
)

type IService interface {
	Login(input InputLoginAdmin) (Admin, error)
	GetUserByID(userID int) (Admin, error)
	GetTotalStatus() (totalPending int64, totalApprove int64)
}

type Service struct {
	repoAdmin       IRepository
	serviceApproval approval.IService
}

func NewService(repoAdmin IRepository, serviceApproval approval.IService) *Service {
	return &Service{repoAdmin, serviceApproval}
}

func (s *Service) Login(input InputLoginAdmin) (Admin, error) {
	// check admin by email
	adminByEmail, err := s.repoAdmin.FindByEmail(input.Email)
	if err != nil {
		return adminByEmail, err
	}

	if adminByEmail.ID == 0 {
		return adminByEmail, fmt.Errorf(
			"email %v tidak ditemukan",
			input.Email,
		)
	}

	// check password
	if err := helper.VerifyPassword(input.Password, adminByEmail.Password); err != nil {
		return adminByEmail, fmt.Errorf("password salah")
	}

	return adminByEmail, nil
}

func (s *Service) GetUserByID(userID int) (Admin, error) {
	adminByID, err := s.repoAdmin.FindByID(userID)
	if err != nil {
		return adminByID, err
	}

	if adminByID.ID == 0 {
		return adminByID, fmt.Errorf(
			"user not found with id %v",
			userID,
		)
	}

	return adminByID, nil
}

func (s *Service) GetTotalStatus() (totalPending int64, totalApprove int64) {
	totalPending, totalApprove = s.serviceApproval.GetTotalStatus()
	return totalPending, totalApprove
}
