package repo

import (
	"context"
	"gorm.io/gorm"
	"plc_auth/domain/model"
)

type PLKRepo struct {
	db *gorm.DB
}

func NewPLkRepo(db *gorm.DB) *PLKRepo {
	return &PLKRepo{db}
}

func (d *PLKRepo) GetKeyByCompany(ctx context.Context, company string) *model.PlcKey {
	var plcKey = model.PlcKey{}
	d.db.Where("industry_id = ?", company).Find(&plcKey)
	return &plcKey
}
