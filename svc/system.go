package svc

import (
	"core/repository"
	"core/serializers"
)

type ISystem interface {
	GetHealth() (*serializers.HealthResp, error)
}

type system struct {
	repo repository.ISystem
}

func NewSystemService(sysrepo repository.ISystem) ISystem {
	return &system{
		repo: sysrepo,
	}
}

func (sys *system) GetHealth() (*serializers.HealthResp, error) {
	resp := serializers.HealthResp{}

	// check db
	dbOnline, err := sys.repo.DBCheck()
	resp.DBOnline = dbOnline

	if err != nil {
		return &resp, err
	}

	return &resp, nil
}
