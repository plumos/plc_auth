package model

type PlcKey struct {
	Id          uint64 `gorm:"column:id;type:int(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Public      string `gorm:"column:public;type:varchar(256);NOT NULL" json:"public"`
	Private     string `gorm:"column:private;type:varchar(256);NOT NULL" json:"private"`
	IndustryId  string `gorm:"column:industry_id;type:varchar(256);NOT NULL" json:"industry_id"`
	CompanyName string `gorm:"column:company_name;type:varchar(256);NOT NULL" json:"company_name"`
	Ctime       int64  `gorm:"column:ctime;type:varchar(256);NOT NULL" json:"ctime"`
	Mtime       int64  `gorm:"column:mtime;type:varchar(256);NOT NULL" json:"mtime"`
}

func (k PlcKey) Table() string {
	return "plc_key"
}
