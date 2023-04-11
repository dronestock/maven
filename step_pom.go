package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/goexl/gox"
	"github.com/goexl/gox/rand"
)

const (
	keyProject = "project"

	keyEnabled              = "enabled"
	keyUpdatePolicy         = "updatePolicy"
	xmlCentral              = "central"
	xmlCentralUrl           = "http://central"
	centralRepositoryFormat = "repository[url='%s']"
	centralPluginFormat     = "pluginRepository[url='%s']"

	repositoryFormat         = "repository[id='%s']"
	snapshotRepositoryFormat = "snapshotRepository[id='%s']"
	keyDistribution          = "distributionManagement"
)

type stepPom struct {
	*plugin
}

func newPomStep(plugin *plugin) *stepPom {
	return &stepPom{
		plugin: plugin,
	}
}

func (p *stepPom) Runnable() bool {
	return 0 != len(p.Repositories)
}

func (p *stepPom) Run(_ context.Context) (err error) {
	original := filepath.Join(p.Source, pomFilename)
	doc := etree.NewDocument()
	if err = doc.ReadFromFile(original); nil != err {
		return
	}

	// 设置项目
	project := p.writeProject(doc)
	// 设置项目属性
	p.writeProperties(project)
	// 设置仓库
	p.writeRepositories(project)
	// 设置发布仓库
	p.writeDistribution(project)
	// 设置发布插件
	p.writePlugins(project)

	// 写入文件
	doc.Indent(xmlSpaces)
	p.pom = gox.StringBuilder(rand.New().String().Length(randLength).Build().Generate(), dot, pomFilename).String()
	if err = doc.WriteToFile(p.pom); nil == err {
		p.Cleanup().Name("清理模块文件").File(p.pom).Build()
	}

	return
}

func (p *stepPom) writeProject(doc *etree.Document) (project *etree.Element) {
	project = doc.SelectElement(keyProject)
	if nil == project {
		doc.CreateProcInst(keyXml, xmlDeclare)
		project = doc.CreateElement(keyProject)
		project.CreateAttr(keyXmlns, xmlProjectXmlns)
		project.CreateAttr(keyXsi, xmlProjectXsi)
		project.CreateAttr(keySchema, xmlProjectSchema)
	}

	group := project.SelectElement(keyGroupId)
	if nil == group {
		group = project.CreateElement(keyGroupId)
	}
	if "" != p.Group {
		group.SetText(p.Group)
	}

	artifact := project.SelectElement(keyArtifactId)
	if nil == artifact {
		artifact = project.CreateElement(keyArtifactId)
	}
	if "" != p.Artifact {
		artifact.SetText(p.Artifact)
	}

	version := project.SelectElement(keyVersion)
	if nil == version {
		version = project.CreateElement(keyVersion)
	}
	if "" != p.Version {
		version.SetText(p.Version)
	}

	packaging := project.SelectElement(keyPackaging)
	if nil == packaging {
		packaging = project.CreateElement(keyPackaging)
	}
	if "" != p.Packaging {
		packaging.SetText(p.Packaging)
	}

	return
}

func (p *stepPom) writeProperties(project *etree.Element) {
	properties := project.SelectElement(keyProperties)
	if nil == properties {
		properties = project.CreateElement(keyProperties)
	}

	for key, value := range p.properties() {
		property := properties.SelectElement(key)
		if nil == property {
			property = properties.CreateElement(key)
		}
		property.SetText(value)
	}
}

func (p *stepPom) writeRepositories(project *etree.Element) {
	repositories := project.SelectElement(keyRepositories)
	if nil == repositories {
		repositories = project.CreateElement(keyRepositories)
	}

	_repository := repositories.FindElementPath(etree.MustCompilePath(fmt.Sprintf(centralRepositoryFormat, xmlCentralUrl)))
	if nil == _repository {
		_repository = repositories.CreateElement(keyRepository)
		_repository.CreateElement(keyId).SetText(xmlCentral)
		_repository.CreateElement(keyUrl).SetText(xmlCentralUrl)
		_repository.CreateElement(keyReleases).CreateElement(keyEnabled).SetText(xmlTrue)
		_repository.CreateElement(keySnapshots).CreateElement(keyEnabled).SetText(xmlTrue)
	}
	// 写入镜像仓库，解决一部分包不能从私有仓库下载的问题
	for _, repo := range p.Repositories {
		// 写入正式仓库
		p.writeRepository(repositories, repo.releaseId(), repo.Release, repositoryFormat, keyRepository)
		// 写入快照仓库
		p.writeRepository(repositories, repo.snapshotId(), repo.Snapshot, repositoryFormat, keyRepository)
	}

	plugins := project.SelectElement(keyPluginRepositories)
	if nil == plugins {
		plugins = project.CreateElement(keyPluginRepositories)
	}

	_plugin := plugins.FindElementPath(etree.MustCompilePath(fmt.Sprintf(centralPluginFormat, xmlCentralUrl)))
	if nil == _plugin {
		_plugin = plugins.CreateElement(keyPluginRepository)
		_plugin.CreateElement(keyId).SetText(xmlCentral)
		_plugin.CreateElement(keyUrl).SetText(xmlCentralUrl)
		_plugin.CreateElement(keyReleases).CreateElement(keyEnabled).SetText(xmlTrue)

		snapshots := _plugin.CreateElement(keySnapshots)
		snapshots.CreateElement(keyEnabled).SetText(xmlTrue)
		snapshots.CreateElement(keyUpdatePolicy).SetText(xmlAlways)
	}
}

func (p *stepPom) writeDistribution(project *etree.Element) {
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	for _, repo := range p.Repositories {
		// 写入正式仓库
		p.writeRepository(distribution, repo.releaseId(), repo.Release, repositoryFormat, keyRepository)
		// 写入快照仓库
		p.writeRepository(distribution, repo.snapshotId(), repo.Snapshot, snapshotRepositoryFormat, keySnapshotRepository)
	}
}

func (p *stepPom) writeRepository(element *etree.Element, id string, url string, format string, key string) {
	path := etree.MustCompilePath(fmt.Sprintf(format, id))
	repo := element.FindElementPath(path)
	if nil != repo {
		element.RemoveChildAt(repo.Index())
	}
	repo = element.CreateElement(key)
	if repo.Text() != id {
		repo.CreateElement(keyId).SetText(id)
		repo.CreateElement(keyName).SetText(id)
		repo.CreateElement(keyUrl).SetText(url)
	}
}
