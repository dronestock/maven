package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type config struct {
	drone.Config

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required_without=Folders"`
	// 目录列表
	Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS}}" validate:"required_without=Folder"`
}

func (c *config) Fields() gox.Fields {
	return []gox.Field{
		field.Strings(`folders`, c.Folders...),
	}
}
