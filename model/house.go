package model

// 爬虫数据表
type Houses struct {
	Model
	MasterId       int64    `gorm:"column:master_id;NOT NULL" json:"master_id"`               // 数据在主体网站的标识id
	Title          string   `gorm:"column:title;NOT NULL" json:"title"`                       // 标题
	TotalPrice     string   `gorm:"column:total_price;NOT NULL" json:"total_price"`           // 挂牌总价
	TotalArea      string   `gorm:"column:total_area;NOT NULL" json:"total_area"`             // 建筑面积（平方米）
	UnitPriceValue string   `gorm:"column:unit_price_value;NOT NULL" json:"unit_price_value"` // 挂牌单价/平方米
	CommunityName  string   `gorm:"column:community_name;NOT NULL" json:"community_name"`     // 小区名
	Area           string   `gorm:"column:area;NOT NULL" json:"area"`                         // 区域，例如（海珠区）
	Addr           string   `gorm:"column:addr;NOT NULL" json:"addr"`                         // 详细地址
	DoorModel      string   `gorm:"column:door_model;NOT NULL" json:"door_model"`             // 户型
	HasElevator    int      `gorm:"column:has_elevator;NOT NULL" json:"has_elevator"`         // 是否有电梯。0=没有，1=有，2=不确定
	HasSubway      int      `gorm:"column:has_subway;NOT NULL" json:"has_subway"`             // 是否近地铁。0=没有，1=有，2=不确定
	Toward         string   `gorm:"column:toward;NOT NULL" json:"toward"`                     // 房屋朝向
	Establish      int      `gorm:"column:establish;NOT NULL" json:"establish"`               // 建立年份
	SpiderSrcUrl   string   `gorm:"column:spider_src_url;NOT NULL" json:"spider_src_url"`     // 数据原始url
	SpiderType     int      `gorm:"column:spider_type;NOT NULL" json:"spider_type"`           // 数据来源。1=贝壳，2=安居客
	Images         []Images `gorm:"foreignkey:HouseId" json:"images"`
}
