//go:build gin && windows && embed

package core

var Static = new(static)

type static struct{}

func (c *static) Set() {
	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build, 并把dist文件夹放到resource文件夹下, 再打开下面1行注释
	// global.Engine.Static("admin", "./static/dist")
	// global.Engine.Static("form-generator", "./static/page")
}
