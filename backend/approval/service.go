package approval

import (
	"fmt"
	"net/http"
)

type IService interface {
	CreateApproval(input InputNewApproval) (Approval, int, error)
	GetAllApprovals() ([]Approval, int, error)
	GetTotalStatus() (totalPending int64, totalApprove int64)
	ChangeStatus(ID int, status int) (Approval, int, error)
}

type service struct {
	repoApproval IRespoistory
}

func NewService(repoApproval IRespoistory) *service {
	return &service{repoApproval}
}

func (s *service) CreateApproval(input InputNewApproval) (Approval, int, error) {
	var approval Approval
	approval.JenisBeasiswa = input.JenisBeasiswa
	approval.Status = input.Status
	approval.UserID = input.UserID

	// call service
	newApproval, err := s.repoApproval.Save(approval)
	if err != nil {
		return newApproval, http.StatusBadRequest, err
	}

	return newApproval, http.StatusOK, nil

}

func (s *service) GetAllApprovals() ([]Approval, int, error) {
	approvals, err := s.repoApproval.FindAll()
	if err != nil {
		return approvals, http.StatusInternalServerError, err
	}

	return approvals, http.StatusOK, nil
}

func (s *service) GetTotalStatus() (totalPending int64, totalApprove int64) {
	totalPending, totalApprove = s.repoApproval.CountStatus(0, 1)
	return totalPending, totalApprove
}

func (s *service) ChangeStatus(ID int, status int) (Approval, int, error) {
	approvalByID, err := s.repoApproval.FindByID(ID)
	if err != nil {
		return approvalByID, http.StatusNotFound, err
	}

	// if not exist
	if approvalByID.ID == 0 {
		return approvalByID, http.StatusNotFound, fmt.Errorf(
			"id %v tidak ditemukan",
			ID,
		)
	}

	// bind
	approvalByID.Status = status

	// call repo
	approvalUpdated, err := s.repoApproval.Update(approvalByID)
	if err != nil {
		return approvalByID, http.StatusNotFound, err
	}

	return approvalUpdated, http.StatusOK, nil
}
