package model


// 爬虫数据表
type Houses struct {
	Model
	MasterId       int64     `gorm:"column:master_id;NOT NULL"`        // 数据在主体网站的标识id
	Title          string    `gorm:"column:title;NOT NULL"`            // 标题
	TotalPrice     string    `gorm:"column:total_price;NOT NULL"`      // 挂牌总价
	TotalArea      string    `gorm:"column:total_area;NOT NULL"`       // 建筑面积（平方米）
	UnitPriceValue string    `gorm:"column:unit_price_value;NOT NULL"` // 挂牌单价/平方米
	CommunityName  string    `gorm:"column:community_name;NOT NULL"`   // 小区名
	Area           string    `gorm:"column:area;NOT NULL"`             // 区域，例如（海珠区）
	Addr           string    `gorm:"column:addr;NOT NULL"`             // 详细地址
	DoorModel      string    `gorm:"column:door_model;NOT NULL"`       // 户型
	HasElevator    int       `gorm:"column:has_elevator;NOT NULL"`     // 是否有电梯。0=没有，1=有，2=不确定
	HasSubway      int       `gorm:"column:has_subway;NOT NULL"`       // 是否近地铁。0=没有，1=有，2=不确定
	Toward         string    `gorm:"column:toward;NOT NULL"`           // 房屋朝向
	Establish      int       `gorm:"column:establish;NOT NULL"`        // 建立年份
	SpiderSrcUrl   string    `gorm:"column:spider_src_url;NOT NULL"`   // 数据原始url
	SpiderType     int       `gorm:"column:spider_type;NOT NULL"`      // 数据来源。1=贝壳，2=安居客
}
