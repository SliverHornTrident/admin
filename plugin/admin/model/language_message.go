package model

import "github.com/nicksnyder/go-i18n/v2/i18n"

// LanguageMessage 语言消息
type LanguageMessage struct {
	Id             string `json:"id" gorm:"index:index_language,priority:2;column:id;comment:唯一标识"`
	Tag            string `json:"tag" gorm:"index:index_language,priority:1;column:tag;comment:语言标识"`
	Hash           string `json:"hash" gorm:"column:hash;comment:哈希"`
	Description    string `json:"description" gorm:"type:text;column:description;comment:文本"`
	LeftDelimiter  string `json:"leftDelimiter" gorm:"column:left_delimiter;comment:模版左标记"`
	RightDelimiter string `json:"rightDelimiter" gorm:"column:right_delimiter;comment:模版右标识"`
	Zero           string `json:"zero" gorm:"type:text;column:zero;comment:zero消息内容"`
	One            string `json:"one" gorm:"type:text;column:one;comment:one消息内容"`
	Two            string `json:"two" gorm:"type:text;column:two;comment:two消息内容"`
	Few            string `json:"few" gorm:"type:text;column:few;comment:few消息内容"`
	Many           string `json:"many" gorm:"type:text;column:many;comment:many消息内容"`
	Other          string `json:"other" gorm:"type:text;column:other;comment:other消息内容"`
}

func (m *LanguageMessage) Message() *i18n.Message {
	return &i18n.Message{
		ID:          m.Id,
		Hash:        m.Hash,
		Description: m.Description,
		LeftDelim:   m.LeftDelimiter,
		RightDelim:  m.RightDelimiter,
		Zero:        m.Zero,
		One:         m.One,
		Two:         m.Two,
		Few:         m.Few,
		Many:        m.Many,
		Other:       m.Other,
	}
}

func (m *LanguageMessage) TableName() string {
	return "shadow_language_messages"
}
