package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
	model2 "github.com/SliverHornTrident/shadow/plugin/admin/model"
	"github.com/spf13/viper"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (c *_gorm) Name() string {
	return "gorm"
}

func (c *_gorm) Viper(viper *viper.Viper) error {
	return nil
}

func (c *_gorm) IsPanic() bool {
	return true
}

func (c *_gorm) ConfigName() string {
	return ""
}

func (c *_gorm) Initialization(ctx context.Context) error {
	models := []any{
		new(model2.Api),
		new(model2.Dictionary),
		new(model2.DictionaryDetail),
		new(model2.Language),
		new(model2.LanguageMessage),
		new(model2.Menu),
		new(model2.MenuButton),
		new(model2.MenuButtonsApis),
		new(model2.MenuParameter),
		new(model2.MenusApis),
		new(model2.Role),
		new(model2.RolesMenus),
		new(model2.RolesMenuButtons),
		new(model2.User),
		new(model2.UsersRoles),
	}
	err := global.Mysql.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
