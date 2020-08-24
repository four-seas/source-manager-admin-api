package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Model struct {
	ID        uint     `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt DateTime `gorm:"default:CURRENT_TIMESTAMP;type:datetime;NOT NULL" json:"created_at"`
	UpdatedAt DateTime `gorm:"default:CURRENT_TIMESTAMP;type:datetime;NOT NULL" json:"updated_at"`
}

type DateTime struct {
	time.Time
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))

	return []byte(formatted), nil
}

func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return t.Time, nil
}

func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
