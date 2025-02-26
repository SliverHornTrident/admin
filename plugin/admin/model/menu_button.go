package model

import "github.com/SliverHornTrident/shadow/global"

// MenuButton 菜单按钮
type MenuButton struct {
	global.Model
	Icon   string `json:"icon" gorm:"column:icon;comment:图标"`
	Name   string `json:"name" gorm:"column:name;comment:按钮名称"`
	MenuId uint   `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	Menu   *Menu  `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
	Apis   []*Api `json:"apis" gorm:"many2many:menu_buttons_apis;foreignKey:ID;joinForeignKey:MenuButtonId;References:ID;JoinReferences:ApiId"`
}

func (m *MenuButton) TableName() string {
	return "shadow_menu_buttons"
}
