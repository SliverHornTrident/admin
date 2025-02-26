package model

import "github.com/SliverHornTrident/shadow/global"

type Dictionary struct {
	global.Model
	Key         string              `json:"key" gorm:"column:key;comment:键"`
	Type        string              `json:"type" gorm:"column:type;comment:类型"`
	Value       string              `json:"value" gorm:"column:value;comment:值"`
	Sort        int                 `json:"sort" gorm:"column:sort;comment:排序"`
	Status      bool                `json:"status" gorm:"column:status;comment:状态"`
	Description string              `json:"desc" gorm:"column:desc;comment:描述"`
	Details     []*DictionaryDetail `json:"details" gorm:"foreignKey:DictionaryId;references:ID"`
}

func (Dictionary) TableName() string {
	return "shadow_dictionaries"
}
