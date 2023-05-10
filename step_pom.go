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

	releaseFormat  = "repository[url='%s']"
	snapshotFormat = "snapshotRepository[url='%s']"
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
	return p.override()
}

func (p *stepPom) Run(_ context.Context) (err error) {
	// 设置项目
	project := p.writeProject()
	// 设置项目属性
	p.writeProperties(project)
	// 设置仓库
	p.writeRepositories(project)
	if p.deploy() { // 只有在需要部署的时候才设置相关配置
		// 设置发布仓库
		p.writeDistribution(project)
		// 设置发布插件
		p.writePlugins(project)
	}

	// 写入文件
	p.pom.Indent(xmlSpaces)
	filename := gox.StringBuilder(rand.New().String().Length(randLength).Build().Generate(), dot, pomFilename).String()
	p.filename = filepath.Join(p.source, filename)
	p.Cleanup().Name("清理模块文件").File(p.filename).Build()
	err = p.pom.WriteToFile(p.filename)

	return
}

func (p *stepPom) writeProject() (project *etree.Element) {
	project = p.pom.SelectElement(keyProject)
	if nil == project {
		p.pom.CreateProcInst(keyXml, xmlDeclare)
		project = p.pom.CreateElement(keyProject)
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
	keyRepositories := "repositories"
	repositories := project.SelectElement(keyRepositories)
	if nil == repositories {
		repositories = project.CreateElement(keyRepositories)
	}

	// 写入镜像仓库，解决一部分包不能从私有仓库下载的问题
	for _, repo := range p.Repositories {
		// 写入正式仓库
		p.writeRepository(repositories, repo.releaseId(p.pom), repo.Release, releaseFormat, keyRepository)
		// 写入快照仓库
		p.writeRepository(repositories, repo.snapshotId(p.pom), repo.Snapshot, releaseFormat, keyRepository)
	}
}

func (p *stepPom) writeDistribution(project *etree.Element) {
	keyDistribution := "distributionManagement"
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	for _, repo := range p.Repositories {
		// 写入正式仓库
		p.writeRepository(distribution, repo.releaseId(p.pom), repo.Release, releaseFormat, keyRepository)
		// 写入快照仓库
		p.writeRepository(distribution, repo.snapshotId(p.pom), repo.Snapshot, snapshotFormat, keySnapshotRepository)
	}
}

func (p *stepPom) writeRepository(element *etree.Element, id string, url string, format string, key string) {
	path := fmt.Sprintf(format, url)
	repo := element.FindElementPath(etree.MustCompilePath(path))
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
