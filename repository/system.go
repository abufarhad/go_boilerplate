package repository

import (
	"gorm.io/gorm"
)

type ISystem interface {
	DBCheck() (bool, error)
}

type system struct {
	*gorm.DB
}

func NewSystemRepository(db *gorm.DB) ISystem {
	return &system{
		DB: db,
	}
}

func (sys *system) DBCheck() (bool, error) {
	dB, _ := sys.DB.DB()
	if err := dB.Ping(); err != nil {
		return false, err
	}

	return true, nil
}
