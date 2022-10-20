package main

import (
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type plugin struct {
	drone.Base

	// 源文件目录
	Source string `default:"${PLUGIN_SOURCE=${SOURCE=.}}" validate:"required"`

	// 仓库
	Repository *repository `default:"${PLUGIN_REPOSITORY=${REPOSITORY}}"`
	// 仓库列表
	Repositories []*repository `default:"${PLUGIN_REPOSITORIES=${REPOSITORIES}}"`

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
	Properties map[string]string `default:"${PLUGIN_PROPERTIES=${PROPERTIES}}"`
	// 参数
	Defines map[string]string `default:"${PLUGIN_DEFINES=${DEFINES}}"`

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

	_passphrase string
}

func newPlugin() drone.Plugin {
	return new(plugin)
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
	if nil != p.Repository {
		p.Repositories = append(p.Repositories, p.Repository)
	}
	if nil == p.Defines {
		p.Defines = make(map[string]string)
	}

	return
}

func (p *plugin) Fields() gox.Fields {
	return []gox.Field{
		field.String(`folder`, p.Source),
	}
}

func (p *plugin) passphrase() (passphrase string) {
	passphrase = p._passphrase
	if `` == strings.TrimSpace(passphrase) {
		passphrase = p.Gpg.Passphrase
	}
	if `` == strings.TrimSpace(passphrase) {
		passphrase = gox.RandString(randLength)
	}
	if `` == p._passphrase {
		p._passphrase = passphrase
	}

	return
}

func (p *plugin) mirrors() (mirrors []string) {
	mirrors = make([]string, 0)
	mirrors = append(mirrors, p.Mirrors...)
	if p.Defaults {
		mirrors = append(mirrors, defaultMirrors...)
	}

	return
}

func (p *plugin) properties() (properties map[string]string) {
	properties = p.Properties
	if !p.Defaults {
		return
	}

	if nil == properties {
		properties = make(map[string]string)
	}
	for key, value := range defaultProperties {
		if _, ok := properties[key]; !ok {
			properties[key] = value
		}
	}

	return
}
