package model

type MenuButtonsApis struct {
	ApiId        uint        `json:"apiId" gorm:"column:api_id;comment:权限Id"`
	Api          *Api        `json:"api" gorm:"foreignKey:ApiId;references:ID"`
	MenuId       uint        `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	Menu         *Menu       `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
	MenuButtonId uint        `json:"menuButtonId" gorm:"column:menu_button_id;comment:菜单按钮Id"`
	MenuButton   *MenuButton `json:"menuButton" gorm:"foreignKey:MenuButtonId;references:ID"`
}

func (m *MenuButtonsApis) TableName() string {
	return "shadow_menu_buttons_apis"
}
