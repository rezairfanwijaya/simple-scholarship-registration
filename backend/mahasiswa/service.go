package mahasiswa

import (
	"fmt"
	"net/http"
	"serkom/approval"
)

type IService interface {
	GetByEmail(email string) (Mahasiswa, int, error)
	RegisterBeasiswa(input InputDaftarBeasiswa) (int, error)
	GetAllRegisteredBeasiswa() ([]registeredMahasiswaOnBeasiswa, int, error)
}

type service struct {
	repoMahasiswa   IRespoistory
	serviceApproval approval.IService
}

const PENDING = 0

func NewService(repoMahasiswa IRespoistory, serviceApproval approval.IService) *service {
	return &service{repoMahasiswa, serviceApproval}
}

func (s *service) GetByEmail(email string) (Mahasiswa, int, error) {
	// find mahasiswa by email
	mahasiswaByEmail, err := s.repoMahasiswa.FindByEmail(email)
	if err != nil {
		return mahasiswaByEmail, http.StatusInternalServerError, err
	}

	// if not exist
	if mahasiswaByEmail.ID == 0 {
		return mahasiswaByEmail, http.StatusNotFound, fmt.Errorf(
			"email %v tidak ditemukan",
			email,
		)
	}

	return mahasiswaByEmail, http.StatusOK, nil
}

func (s *service) RegisterBeasiswa(input InputDaftarBeasiswa) (int, error) {
	// find mahasiswa by email
	mahasiswaByEmail, err := s.repoMahasiswa.FindByEmail(input.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// if not exist
	if mahasiswaByEmail.ID == 0 {
		return http.StatusNotFound, fmt.Errorf(
			"email %v tidak ditemukan",
			input.Email,
		)
	}

	// update data mahasiswa
	mahasiswaByEmail.BerkasURL = input.BerkasURL
	mahasiswaByEmail.Email = input.Email
	mahasiswaByEmail.IPK = input.IPK
	mahasiswaByEmail.Nama = input.Nama
	mahasiswaByEmail.Semester = input.Semester
	mahasiswaByEmail.Telepon = input.Telepon
	_, err = s.repoMahasiswa.Update(mahasiswaByEmail)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// input for approval
	var approval approval.InputNewApproval
	approval.JenisBeasiswa = input.JenisBeassiwa
	approval.Status = PENDING
	approval.UserID = mahasiswaByEmail.ID

	// call service approval
	_, code, err := s.serviceApproval.CreateApproval(approval)
	if err != nil {
		return code, err
	}

	return http.StatusOK, nil
}

func (s *service) GetAllRegisteredBeasiswa() ([]registeredMahasiswaOnBeasiswa, int, error) {
	approvals, code, err := s.serviceApproval.GetAllApprovals()
	if err != nil {
		return []registeredMahasiswaOnBeasiswa{}, code, err
	}

	// get data mahasiswa based on user id in approvals data
	var result []registeredMahasiswaOnBeasiswa
	for _, data := range approvals {
		mahasiswaById, err := s.repoMahasiswa.FindByID(data.UserID)
		if err != nil {
			return []registeredMahasiswaOnBeasiswa{}, http.StatusBadRequest, err
		}

		if mahasiswaById.ID == 0 {
			return []registeredMahasiswaOnBeasiswa{}, http.StatusNotFound, fmt.Errorf(
				"id %v tidak ditemukan",
				data.UserID,
			)
		}

		registeredMahasiswa := registeredMahasiswaOnBeasiswa{
			ID:            data.ID,
			Nama:          mahasiswaById.Nama,
			Email:         mahasiswaById.Email,
			Telepon:       mahasiswaById.Telepon,
			Semester:      mahasiswaById.Semester,
			IPK:           mahasiswaById.IPK,
			JenisBeasiswa: data.JenisBeasiswa,
			BerkasURL:     mahasiswaById.BerkasURL,
			Status:        data.Status,
		}
		// append
		result = append(result, registeredMahasiswa)

	}

	return result, http.StatusOK, nil
}
