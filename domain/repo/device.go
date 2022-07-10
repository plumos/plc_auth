package repo

import (
	"context"
	"gorm.io/gorm"
	"plc_auth/domain/model"
)

type DeviceRepo struct {
	db *gorm.DB
}

func NewDeviceRepo(db *gorm.DB) *DeviceRepo {
	return &DeviceRepo{db}
}

func (d *DeviceRepo) GetDevice(ctx context.Context, code uint64) *model.Device {
	var device = model.Device{}
	d.db.Where("device_code = ?", code).Find(&device)
	return &device
}
