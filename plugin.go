package main

import (
	"strings"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
)

type plugin struct {
	drone.Base

	// 执行程序
	Binary binary `default:"${BINARY}"`
	// 文件路径
	Filepath _filepath `default:"${FILEPATH}"`
	// 源文件目录
	Source string `default:"${SOURCE=.}" validate:"required"`

	// 仓库
	Repository *repository `default:"${REPOSITORY}"`
	// 仓库列表
	Repositories []*repository `default:"${REPOSITORIES}"`

	// 密钥
	Gpg gpg `default:"${GPG}"`

	// 坐标，组
	Group string `default:"${GROUP}"`
	// 坐标，制品
	Artifact string `default:"${ARTIFACT}"`
	// 版本
	Version string `default:"${VERSION}"`
	// 打包方式
	Packaging string `default:"${PACKAGING=jar}" validate:"oneof=jar war"`

	// 额外属性
	Properties map[string]string `default:"${PROPERTIES}"`
	// 参数
	Defines map[string]string `default:"${DEFINES}"`

	// 镜像加速列表
	Mirrors []string `default:"${MIRRORS}"`
	// 测试
	Test bool `default:"${TEST=true}"`
	// 清理
	Clean bool `default:"${CLEAN=true}"`
	// 是否包含源码
	Code *bool `default:"${CODE=true}"`
	// 是否包含文档
	Doc *bool `default:"${DOC=true}"`

	// 打包插件版本
	JarPluginVersion string `default:"${JAR_PLUGIN_VERSION=3.2.1}"`
	// 源码插件版本
	SourcePluginVersion string `default:"${SOURCE_PLUGIN_VERSION=3.2.1}"`
	// 文档插件版本
	DocPluginVersion string `default:"${DOC_PLUGIN_VERSION=3.11.2}"`
	// 签名插件版本
	GpgPluginVersion string `default:"${GPG_PLUGIN_VERSION=3.2.7}"`
	// 发布仓库版本
	CentralPluginVersion string `default:"${NEXUS_PLUGIN_VERSION=0.6.0}"`
	NexusPluginVersion   string `default:"${NEXUS_PLUGIN_VERSION=0.6.0}"`
	// 执行程序
	Java java `default:"${JAVA}" json:"java,omitempty"`

	phrase string
}

func newPlugin() drone.Plugin {
	return new(plugin)
}

func (p *plugin) Config() drone.Config {
	return p
}

func (p *plugin) Steps() drone.Steps {
	return drone.Steps{
		// 执行出错具有不可重复性，不需要重试
		drone.NewStep(newGlobalStep(p)).Name("全局配置").Interrupt().Build(),
		// 执行出错具有不可重复性，不需要重试
		drone.NewStep(newPomStep(p)).Name("项目配置").Interrupt().Build(),
		// 执行出错具有不可重复性，不需要重试
		drone.NewStep(newKeypairStep(p)).Name("生成密钥").Interrupt().Build(),
		drone.NewStep(newPackageStep(p)).Name("编译打包").Build(),
		drone.NewStep(newGskStep(p)).Name("上传密钥").Build(),
		drone.NewStep(newDeployStep(p)).Name("发布仓库").Build(),
	}
}

func (p *plugin) Setup() (err error) {
	if nil != p.Repository {
		p.Repositories = append(p.Repositories, p.Repository)
	}

	return
}

func (p *plugin) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("source", p.Source),
		field.New("version", p.Version),
	}
}

func (p *plugin) passphrase() (passphrase string) {
	passphrase = p.phrase
	if "" == strings.TrimSpace(passphrase) {
		passphrase = p.Gpg.Passphrase
	}
	if "" == strings.TrimSpace(passphrase) {
		passphrase = rand.New().String().Length(randLength).Build().Generate()
	}
	if "" == p.phrase {
		p.phrase = passphrase
	}

	return
}

func (p *plugin) mirrors() (mirrors []string) {
	mirrors = make([]string, 0)
	mirrors = append(mirrors, p.Mirrors...)
	if *p.Defaults {
		mirrors = append(mirrors, defaultMirrors...)
	}

	return
}

func (p *plugin) properties() (properties map[string]string) {
	properties = p.Properties
	if !*p.Defaults {
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

func (p *plugin) defines() (defines map[string]string) {
	defines = p.Defines
	if !*p.Defaults {
		return
	}

	if nil == defines {
		defines = make(map[string]string)
	}
	for key, value := range defaultDefines {
		if _, ok := defines[key]; !ok {
			defines[key] = value
		}
	}

	return
}

func (p *plugin) private() (private bool) {
	private = true
	for _, repo := range p.Repositories {
		private = private && repo.Private
	}

	return
}

func (p *plugin) mirrorOf() string {
	mirrorOf := gox.StringBuilder()
	for _, repo := range p.Repositories {
		mirrorOf.Append(comma).Append(exclamation).Append(repo.id())
	}

	return mirrorOf.String()
}

func (p *plugin) deploy() (deploy bool) {
	for _, _repository := range p.Repositories {
		deploy = nil != _repository.Deploy && *_repository.Deploy
		if deploy {
			break
		}
	}

	return
}

func (p *plugin) override() bool {
	return 0 != len(p.Properties) || 0 != len(p.Defines) || p.deploy()
}

func (p *plugin) code() bool {
	return nil != p.Code && *p.Code
}

func (p *plugin) doc() bool {
	return nil != p.Doc && *p.Doc
}
