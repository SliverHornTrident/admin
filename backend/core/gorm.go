//go:build gorm || gorms

package core

import "gorm.io/gorm"

var Gorm = new(_gorm)

type _gorm struct{}

// Copy dst to src
// if dst is nil, return
// if src is nil, src = dst and return
// if dst is not nil and src is not nil, src = dst
func (g *_gorm) Copy(src **gorm.DB, dst *gorm.DB) {
	if dst == nil {
		return
	}
	if *src == nil {
		*src = dst
		return
	}
	**src = *dst
}
