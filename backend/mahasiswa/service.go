package mahasiswa

import (
	"fmt"
	"net/http"
)

type IService interface {
	GetByEmail(email string) (Mahasiswa, int, error)
}

type service struct {
	repoMahasiswa IRespoistory
}

func NewService(repoMahasiswa IRespoistory) *service {
	return &service{repoMahasiswa}
}

func (s *service) GetByEmail(email string) (Mahasiswa, int, error) {
	mahasiswaByEmail, err := s.repoMahasiswa.FindByEmail(email)
	if err != nil {
		return mahasiswaByEmail, http.StatusInternalServerError, err
	}

	if mahasiswaByEmail.ID == 0 {
		return mahasiswaByEmail, http.StatusNotFound, fmt.Errorf(
			"email %v tidak ditemukan",
			email,
		)
	}

	return mahasiswaByEmail, http.StatusOK, nil
}
