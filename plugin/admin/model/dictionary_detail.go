package model

import (
	"github.com/SliverHornTrident/shadow/global"
)

// DictionaryDetail 字典详情
type DictionaryDetail struct {
	global.Model
	Key          string `json:"key" gorm:"column:key;comment:键"`
	Value        string `json:"value" gorm:"column:value;comment:值"`
	Extend       string `json:"extend" gorm:"column:extend;comment:扩展值"`
	Sort         int    `json:"sort" gorm:"column:sort;comment:排序"`
	Status       bool   `json:"status" gorm:"column:status;comment:状态"`
	DictionaryId uint   `json:"dictionaryId" gorm:"column:dictionary_id;comment:字典id"`
}

func (DictionaryDetail) TableName() string {
	return "shadow_dictionary_details"
}
