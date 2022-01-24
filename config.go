package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`
	// 目录列表
	Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS}}"`
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, c.Folder),
	}
}
