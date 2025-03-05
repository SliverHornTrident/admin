//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
	"strings"
)

func Merge() error {
	output, err := sh.Output("git", "remote")
	if err != nil {
		return errors.Wrap(err, "git remote failed!")
	}
	outputs := strings.Split(output, "\n")
	var has bool
	for i := 0; i < len(outputs); i++ {
		if outputs[i] == "shadow" {
			has = true
		}
	}
	if !has {
		output, err = sh.Output("git", "remote", "add", "shadow", `https://github.com/SliverHornTrident/shadow.git`)
		if err != nil {
			return errors.Wrap(err, "git remote add shadow failed!")
		}
	}
	output, err = sh.Output("git", "fetch", "shadow")
	if err != nil {
		return errors.Wrap(err, "git fetch shadow failed!")
	}
	output, err = sh.Output("git", "merge", "shadow/main")
	if err != nil {
		return errors.Wrap(err, "git merge shadow/main failed!")
	}
	return nil
}
