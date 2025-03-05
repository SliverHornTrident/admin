package model

import "github.com/SliverHornTrident/shadow/global"

// User 用户
type User struct {
	global.Model
	Email    string  `json:"email" gorm:"column:email;comment:邮箱"`
	Phone    string  `json:"phone" gorm:"column:phone;comment:手机号"`
	Avatar   string  `json:"avatar" gorm:"column:avatar;comment:头像"`
	Nickname string  `json:"nickname" gorm:"column:nickname;comment:昵称"`
	Username string  `json:"username" gorm:"column:username;comment:用户名"`
	Password string  `json:"password" gorm:"column:password;comment:密码"`
	RoleId   uint    `json:"roleId" gorm:"column:role_id;comment:活跃角色Id"`
	Role     *Role   `json:"role" gorm:"foreignKey:RoleId;references:ID"`
	Roles    []*Role `json:"roles" gorm:"many2many:users_roles;foreignKey:ID;joinForeignKey:UserId;References:ID;JoinReferences:RoleId"`
}

func (m *User) TableName() string {
	return "shadow_users"
}
