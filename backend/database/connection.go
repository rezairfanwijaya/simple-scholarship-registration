package database

import (
	"fmt"

	"serkom/admin"
	"serkom/helper"
	"serkom/mahasiswa"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(path string) (*gorm.DB, error) {
	// get env file
	env, err := helper.GetENV(path)
	if err != nil {
		return nil, err
	}

	dbUsername := env["DATABASE_USERNAME"]
	dbPassword := env["DATABASE_PASSWORD"]
	dbHost := env["DATABASE_HOST"]
	dbPort := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	// generate dsn
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// open connection to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("failed connect to database :%v", err.Error())
	}

	// migration schema
	if err := db.AutoMigrate(&mahasiswa.Mahasiswa{}, &admin.Admin{}); err != nil {
		return db, fmt.Errorf("failed migrate schema : %v", err.Error())
	}

	// migration admin account
	if err := MigarationAdmin(db); err != nil {
		return db, fmt.Errorf("failed migrate admin account : %v", err.Error())
	}

	// migration mahasiswa account
	if err := MigarationMahasiswa(db); err != nil {
		return db, fmt.Errorf("failed migrate mahasiswa account : %v", err.Error())
	}

	return db, nil
}
