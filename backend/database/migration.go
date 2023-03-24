package database

import (
	"fmt"
	"serkom/admin"
	"serkom/helper"
	"serkom/mahasiswa"

	"gorm.io/gorm"
)

func MigarationAdmin(db *gorm.DB) error {
	// check is admin exist
	exist, err := isExist(db, "admin")
	if err != nil {
		return fmt.Errorf("failed get data admin on migration : %v", err.Error())
	}

	// migration
	if !exist {
		password, err := helper.HashPassowrd("12345")
		if err != nil {
			return fmt.Errorf(
				"failed hash password when migration : %v",
				err.Error(),
			)
		}

		if err := db.Create(&admin.Admin{
			Email:    "admin@gmail.com",
			Password: password,
			Role:     "admin",
		}).Error; err != nil {
			return fmt.Errorf(
				"failed store data admin : %v",
				err.Error(),
			)
		}

		return nil
	}

	return nil
}
func MigarationMahasiswa(db *gorm.DB) error {
	// check is mahasiswa exist
	exist, err := isExist(db, "mahasiswa")
	if err != nil {
		return fmt.Errorf("failed get data mahasiswa on migration : %v", err.Error())
	}

	mahasiswaSeeds := mahasiswa.CreateSeedMahasiswa()

	// migration
	if !exist {
		for _, mahasiswaSeed := range mahasiswaSeeds {
			if err := db.Create(&mahasiswaSeed).Error; err != nil {
				return fmt.Errorf(
					"gagal migration seed mahasiswa : %v",
					err.Error(),
				)
			}
		}
	}

	return nil
}

func isExist(db *gorm.DB, table string) (bool, error) {
	// cek ke tabel admin
	if table == "admin" {

		var admin []admin.Admin

		if err := db.Find(&admin).Error; err != nil {
			return false, err
		}

		if len(admin) == 0 {
			return false, nil
		}

		return true, nil
	}

	// cek ke tabel mahasiswa
	var mahasiswas []mahasiswa.Mahasiswa

	if err := db.Find(&mahasiswas).Error; err != nil {
		return false, err
	}

	if len(mahasiswas) == 0 {
		return false, nil
	}

	return true, nil
}
