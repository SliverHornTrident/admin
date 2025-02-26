package model

type RolesMenuButtons struct {
	RoleId       uint        `json:"roleId" gorm:"column:role_id;comment:角色Id"`
	Role         *Role       `json:"role" gorm:"foreignKey:RoleId;references:ID"`
	MenuId       uint        `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	Menu         *Menu       `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
	MenuButtonId uint        `json:"menuButtonId" gorm:"column:menu_button_id;comment:菜单按钮Id"`
	MenuButton   *MenuButton `json:"menuButton" gorm:"foreignKey:MenuButtonId;references:ID"`
}

func (m *RolesMenuButtons) TableName() string {
	return "shadow_roles_menu_buttons"
}
