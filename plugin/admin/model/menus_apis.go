package model

type MenusApis struct {
	ApiId  uint  `json:"apiId" gorm:"column:api_id;comment:权限Id"`
	Api    *Api  `json:"api" gorm:"foreignKey:ApiId;references:ID"`
	MenuId uint  `json:"menuId" gorm:"column:menu_id;comment:菜单Id"`
	Menu   *Menu `json:"menu" gorm:"foreignKey:MenuId;references:ID"`
}

func (m *MenusApis) TableName() string {
	return "shadow_menus_apis"
}
