package model

type RolesMenus struct {
	RoleId uint  `json:"roleId" gorm:"column:role_id;comment:角色Id"`
	Role   *Role `json:"role" gorm:"foreignKey:RoleId;references:ID"`
	MenuId uint  `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	Menu   *Menu `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
}

func (m *RolesMenus) TableName() string {
	return "shadow_roles_menus"
}
