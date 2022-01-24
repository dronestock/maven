package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.Base

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required_without=Folders"`
	// 目录列表
	Folders []string `default:"${PLUGIN_FOLDERS=${FOLDERS}}" validate:"required_without=Folder"`
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.test, drone.Name(`测试`)),
	}
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.Strings(`folders`, p.Folders...),
	}
}
