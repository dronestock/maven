package main

import (
	`net/url`

	`github.com/dronestock/drone`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

type plugin struct {
	drone.PluginBase

	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}" validate:"required"`

	// 仓库
	Repository repository `default:"${PLUGIN_REPOSITORY=${REPOSITORY}}"`
	// 用户名
	Username string `default:"${PLUGIN_USERNAME=${USERNAME}}"`
	// 密码
	Password string `default:"${PLUGIN_PASSWORD=${PASSWORD}}"`

	// 密钥
	Gpg gpg `default:"${PLUGIN_GPG=${GPG}}"`

	// 坐标，组
	Group string `default:"${PLUGIN_GROUP=${GROUP}}"`
	// 坐标，制品
	Artifact string `default:"${PLUGIN_ARTIFACT=${ARTIFACT}}"`
	// 版本
	Version string `default:"${PLUGIN_VERSION=${VERSION}}"`
	// 打包方式
	Packaging string `default:"${PLUGIN_PACKAGING=${PACKAGING=jar}}" validate:"oneof=jar war"`

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

	// 打包插件版本
	JarPluginVersion string `default:"${PLUGIN_JAR_PLUGIN_VERSION=${JAR_PLUGIN_VERSION=3.2.1}}"`
	// 源码插件版本
	SourcePluginVersion string `default:"${PLUGIN_SOURCE_PLUGIN_VERSION=${SOURCE_PLUGIN_VERSION=3.2.1}}"`
	// 文档插件版本
	DocPluginVersion string `default:"${PLUGIN_DOC_PLUGIN_VERSION=${DOC_PLUGIN_VERSION=3.3.1}}"`
	// 签名插件版本
	GpgPluginVersion string `default:"${PLUGIN_GPG_PLUGIN_VERSION=${GPG_PLUGIN_VERSION=3.0.1}}"`
	// 发布仓库版本
	NexusPluginVersion string `default:"${PLUGIN_NEXUS_PLUGIN_VERSION=${NEXUS_PLUGIN_VERSION=1.6.3}}"`

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
		drone.NewStep(p.keypair, drone.Name(`生成密钥`)),
		drone.NewStep(p.global, drone.Name(`写入全局配置`)),
		drone.NewStep(p.pom, drone.Name(`修改项目配置`)),
		drone.NewStep(p.pkg, drone.Name(`打包`)),
		drone.NewStep(p.gsk, drone.Name(`上传密钥到服务器`)),
		drone.NewStep(p.deploy, drone.Name(`发布到仓库`)),
	}
}

func (p *plugin) Setup() (unset bool, err error) {
	p.Parse(p.__properties, p.Properties...)

	return
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Source),
	}
}

func (p *plugin) repositoryId(link string) (id string) {
	if uri, err := url.Parse(link); nil != err {
		id = gox.RandString(randLength)
	} else {
		id = uri.Host
	}

	return
}

func (p *plugin) _mirrors() (mirrors []string) {
	mirrors = make([]string, 0)
	mirrors = append(mirrors, p.Mirrors...)
	if p.Defaults {
		mirrors = append(mirrors, defaultMirrors...)
	}

	return
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
