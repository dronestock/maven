package main

import (
	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 目录
	Folder string `default:"${PLUGIN_FOLDER=${FOLDER=.}}" validate:"required"`

	// 正式仓库
	// nolint:lll
	Repository string `default:"${PLUGIN_REPOSITORY=${REPOSITORY=https://oss.sonatype.org/service/local/staging/deploy/maven2}}"`
	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${USERNAME}}"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${PASSWORD}}"`

	// 坐标，组
	Group string `default:"${PLUGIN_GROUP=${GROUP}}"`
	// 坐标，制品
	Artifact string `default:"${PLUGIN_ARTIFACT=${ARTIFACT}}"`
	// 版本
	Version string `default:"${PLUGIN_VERSION=${VERSION}}"`
	// 打包方式
	Packaging string `default:"${PLUGIN_PACKAGING=${PACKAGING}}"`

	// 额外属性
	Properties []string `default:"${PLUGIN_PROPERTIES=${PROPERTIES}}"`

	// 镜像加速列表
	Mirrors []string `default:"${PLUGIN_MIRRORS=${MIRRORS}}"`
	// 测试
	Test bool `default:"${PLUGIN_TEST=${TEST=true}}"`
	// 清理
	Clean bool `default:"${PLUGIN_CLEAN=${CLEAN=true}}"`
	// 是否包含源码
	Sources bool `default:"${PLUGIN_SOURCES=${SOURCES=true}}"`
	// 是否包含文档
	Docs bool `default:"${PLUGIN_DOCS=${DOCS=true}}"`

	__properties map[string]string
}

func newPlugin() drone.Plugin {
	return &plugin{
		__properties: make(map[string]string),
	}
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.settings, drone.Name(`写入全局配置`)),
		drone.NewStep(p.pom, drone.Name(`修改项目配置`), drone.Break()),
		drone.NewStep(p.do, drone.Name(`执行Maven操作`)),
	}
}

func (p *plugin) Setup() (unset bool, err error) {
	p.Parse(p.__properties, p.Properties...)

	return
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Folder),
	}
}

func (p *plugin) _properties() (properties map[string]string) {
	properties = p.__properties
	if !p.Defaults {
		return
	}

	for key, value := range defaultProperties {
		if _, ok := properties[key]; !ok {
			properties[key] = value
		}
	}

	return
}
