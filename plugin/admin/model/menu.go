package model

import "github.com/SliverHornTrident/shadow/global"

// Menu 菜单
type Menu struct {
	global.Model
	Name       string           `json:"name" gorm:"column:name;comment:菜单名称"`
	Icon       string           `json:"icon" gorm:"column:icon;comment:菜单图标"`
	Path       string           `json:"path" gorm:"column:path;comment:菜单路径"`
	Title      string           `json:"title" gorm:"column:title;comment:菜单标题"`
	Component  string           `json:"component" gorm:"column:component;comment:菜单文件路径"`
	Highlight  string           `json:"highlight" gorm:"column:highlight;comment:菜单高亮"`
	IsHidden   bool             `json:"hidden" gorm:"column:hidden;comment:是否隐藏"`
	AutoClose  bool             `json:"autoClose" gorm:"column:auto_close;comment:是否自动关闭"`
	IsDefault  bool             `json:"isDefault" gorm:"column:is_default;comment:是否基础路由"`
	KeepAlive  bool             `json:"keepAlive" gorm:"column:keep_alive;comment:是否缓存"`
	Sort       int              `json:"sort" gorm:"column:sort;comment:菜单排序"`
	Level      uint             `json:"level" gorm:"column:level;comment:菜单层级"`
	ParentId   uint             `json:"parentId" gorm:"column:parent_id;comment:父级菜单Id"`
	Apis       []*Api           `json:"apis" gorm:"many2many:menus_apis;foreignKey:ID;joinForeignKey:MenuId;References:ID;JoinReferences:ApiId"`
	Buttons    []*MenuButton    `json:"buttons" gorm:"foreignKey:MenuId;references:ID"`
	Parameters []*MenuParameter `json:"parameters" gorm:"foreignKey:MenuId;references:ID"`
}

func (m *Menu) TableName() string {
	return "shadow_menus"
}
