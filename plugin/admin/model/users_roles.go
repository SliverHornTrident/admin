package model

type UsersRoles struct {
	RoleId uint  `json:"roleId" gorm:"column:role_id;comment:角色Id"`
	Role   *Role `json:"role" gorm:"foreignKey:RoleId;references:ID"`
	UserId uint  `json:"userId" gorm:"column:user_id;comment:用户Id"`
	User   *User `json:"user" gorm:"foreignKey:UserId;references:ID"`
}

func (m *UsersRoles) TableName() string {
	return "shadow_users_roles"
}
