package core

import (
	"context"
	"github.com/SliverHornTrident/shadow/global"
	"github.com/SliverHornTrident/shadow/plugin/admin/model"
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
		new(model.Api),
		new(model.Dictionary),
		new(model.DictionaryDetail),
		new(model.Language),
		new(model.LanguageMessage),
		new(model.Menu),
		new(model.MenuButton),
		new(model.MenuButtonsApis),
		new(model.MenuParameter),
		new(model.MenusApis),
		new(model.Role),
		new(model.RolesMenus),
		new(model.RolesMenuButtons),
		new(model.User),
		new(model.UsersRoles),
	}
	err := global.Mysql.AutoMigrate(models...)
	if err != nil {
		return err
	}
	return nil
}
