package model

import "gorm.io/gorm"

type SysBaseMenuParameter struct {
	gorm.Model
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"commit:'地址栏携带参数为params还是query'"`
	Key           string `json:"key" gorm:"commit:'地址栏携带参数的key'"`
	Value         string `json:"value" gorm:"commit:'地址栏携带参数的值'"`
}
