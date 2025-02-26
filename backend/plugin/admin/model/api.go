package model

import "github.com/SliverHornTrident/shadow/global"

// Api 接口
type Api struct {
	global.Model
	Path        string `json:"path" gorm:"column:path;comment:路径"`
	Group       string `json:"group" gorm:"column:group;comment:分组"`
	Method      string `json:"method" gorm:"column:method;comment:方法"`
	Description string `json:"description" gorm:"column:description;comment:描述"`
}

func (m *Api) TableName() string {
	return "shadow_apis"
}
