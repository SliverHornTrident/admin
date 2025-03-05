package main

import (
	"github.com/SliverHornTrident/shadow/plugin/admin/model"
	"gorm.io/gen"
	"path/filepath"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Join("plugin", "admin", "model", "dao"),                   // output path
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	g.ApplyBasic(
		new(model.Api),

		new(model.User),
		new(model.UsersRoles),

		new(model.Menu),
		new(model.MenusApis),
		new(model.MenuButton),
		new(model.MenuButtonsApis),
		new(model.MenuParameter),

		new(model.Role),
		new(model.RolesMenus),
		new(model.RolesMenuButtons),

		new(model.Dictionary),
		new(model.DictionaryDetail),

		new(model.Language),
		new(model.LanguageMessage),
	)

	// Generate the code
	g.Execute()
}
