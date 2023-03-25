package approval

import "net/http"

type IService interface {
	CreateApproval(input InputNewApproval) (Approval, int, error)
	GetAllApprovals() ([]Approval, int, error)
	GetTotalStatus() (totalPending int64, totalApprove int64)
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
