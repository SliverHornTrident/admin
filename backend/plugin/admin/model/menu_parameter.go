package model

import "github.com/SliverHornTrident/shadow/global"

// MenuParameter 菜单参数
type MenuParameter struct {
	global.Model
	Key    string `json:"key" gorm:"column:key;comment:地址栏携带参数的键"`
	Type   string `json:"type" gorm:"column:type;comment:地址栏携带参数(params|query)"`
	Value  string `json:"value" gorm:"column:value;comment:地址栏携带参数的值"`
	MenuId uint   `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
}

func (m *MenuParameter) TableName() string {
	return "shadow_menu_parameters"
}
