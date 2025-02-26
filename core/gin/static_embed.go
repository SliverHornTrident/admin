package core

import (
	"embed"
	"io/fs"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

// fsFunc Define redirect able embed fs interface
type fsFunc func(name string) (fs.File, error)

// Open .
func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// StaticFSFromEmbed
// Redirect able embed files
// use :
//
//	//go:embed static
//	var static embed.FS
//	......
//	router.StaticFSFromEmbed("/", "static/", static)
func StaticFSFromEmbed(group *gin.RouterGroup, relativePath, root string, assets embed.FS) gin.IRoutes {
	system := http.FS(fsFunc(func(name string) (fs.File, error) {
		assetPath := path.Join(root, name)
		// If we can't find the asset, system can handle the error
		file, err := assets.Open(assetPath)
		if err != nil {
			return nil, err
		}
		// Otherwise assume this is a legitimate request routed correctly
		return file, err
	}))
	return group.StaticFS(relativePath, system)
}
