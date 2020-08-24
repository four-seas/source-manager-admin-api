package model

// 房子照片信息表
type Images struct {
	Model
	HouseId int    `gorm:"column:house_id;NOT NULL" json:"house_id"`
	Path    string `gorm:"column:path;NOT NULL" json:"path"`
	Url     string `gorm:"column:url;NOT NULL" json:"url"`
}
