package model

type Device struct {
	Id          uint64 `gorm:"column:id;type:int(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	DeviceCode  uint64 `gorm:"column:device_code;type:int(20) unsigned;NOT NULL" json:"device_code"`
	IndustryId  string `gorm:"column:industry_id;type:varchar(256);NOT NULL" json:"industry_id"`
	CompanyName string `gorm:"column:company_name;type:varchar(256);NOT NULL" json:"company_name"`
	Status      int    `gorm:"column:status;type:varchar(256);NOT NULL" json:"status"`

	Ctime int64 `gorm:"column:ctime;type:varchar(256);NOT NULL" json:"ctime"`
	Mtime int64 `gorm:"column:mtime;type:varchar(256);NOT NULL" json:"mtime"`
}

func (k Device) Table() string {
	return "device"
}
