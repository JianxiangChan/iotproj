package model

import (
	"time"
)

type TabAllSensorData struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mac         string    `json:"mac" xorm:"VARCHAR(16)"`
	Temperature float32   `json:"temperature" xorm:"DOUBLE(8,3)"`
	Led         bool      `json:"led" xorm:"BOOL"`
	CreateDate  time.Time `json:"create_date" xorm:"DATETIME"`
}

type TabAllSensorDataLast struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Mac         string    `json:"mac" xorm:"VARCHAR(16)"`
	Temperature float32   `json:"temperature" xorm:"DOUBLE(8,3)"`
	Led         bool      `json:"led" xorm:"BOOL"`
	CreateDate  time.Time `json:"create_date" xorm:"DATETIME"`
}
