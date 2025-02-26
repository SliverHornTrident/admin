package model

import "github.com/SliverHornTrident/shadow/global"

// Role 角色
type Role struct {
	global.Model
	Name        string        `json:"name" gorm:"column:name;comment:角色名称"`
	Menu        string        `json:"menu" gorm:"column:menu;comment:默认菜单"`
	ParentId    uint          `json:"parentId" gorm:"column:parent_id;comment:父级角色Id"`
	Overstep    bool          `json:"overstep" gorm:"column:overstep;comment:越权菜单"`
	Menus       []*Menu       `json:"menus" gorm:"many2many:roles_menus;foreignKey:ID;joinForeignKey:RoleId;References:ID;JoinReferences:MenuId"`
	MenuButtons []*MenuButton `json:"menuButtons" gorm:"many2many:roles_menu_buttons;foreignKey:ID;joinForeignKey:RoleId;References:ID;JoinReferences:MenuButtonId"`
}

func (m *Role) TableName() string {
	return "shadow_roles"
}
