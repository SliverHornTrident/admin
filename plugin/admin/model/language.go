package model

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

// Language 语言表
type Language struct {
	Tag       string             `json:"tag" gorm:"PrimaryKey;index:language_tag_index;column:tag;comment:语言标识"`
	Path      string             `json:"path" gorm:"column:path;comment:路径"`
	Name      string             `json:"name" gorm:"column:name;comment:语言名称"`
	Unmarshal string             `json:"unmarshal" gorm:"column:unmarshal;comment:序列化"`
	Enable    bool               `json:"enable" gorm:"column:enable;comment:启用"`
	Messages  []*LanguageMessage `json:"messages" gorm:"foreignKey:Tag;references:Tag"`
}

func (m *Language) TableName() string {
	return "shadow_languages"
}

func (m *Language) BeforeCreate(tx *gorm.DB) error {
	if m.Path == "" {
		m.Path = strings.Join([]string{m.Tag, m.Unmarshal}, ".")
	}
	return nil
}

func (m *Language) BeforeUpdate(tx *gorm.DB) error {
	var entity Language
	err := tx.WithContext(tx.Statement.Context).Model(m).Where("enable = ?", true).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // 没有启用的语言
		m.Enable = true
	} else {
		if m.Enable {
			entity.Enable = false
			err = tx.WithContext(tx.Statement.Context).Model(m).Save(&entity).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
