package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/beevik/etree"
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

func (p *stepPom) Run(ctx context.Context) (err error) {
	// 写入镜像仓库，解决一部分包不能从私有仓库下载的问题
	for _, repo := range p.Repositories {
		err = p.run(ctx, repo)
		if nil != err {
			break
		}
	}

	return
}

func (p *stepPom) run(_ context.Context, repo *repository) (err error) {
	original := filepath.Join(p.Source, pomFilename)
	pom := etree.NewDocument()
	if rfe := pom.ReadFromFile(original); nil != rfe {
		err = rfe
	} else {
		err = p.writeFile(pom, repo)
	}

	return
}

func (p *stepPom) writeFile(pom *etree.Document, repo *repository) (err error) {
	// 设置项目
	project := p.writeProject(pom)
	// 设置项目属性
	p.writeProperties(project)
	// 设置仓库
	p.writeRepositories(project)
	if p.deploy() { // 只有在需要部署的时候才设置相关配置
		// 设置发布仓库
		p.writeDistribution(project)
		// 设置发布插件
		p.writePlugins(project, repo)
	}

	// 写入文件
	pom.Indent(xmlSpaces)
	filename := repo.filename(p.Source)
	p.Cleanup().Name("清理模块文件").File(filename).Build()
	err = pom.WriteToFile(filename)

	return
}

func (p *stepPom) writeProject(pom *etree.Document) (project *etree.Element) {
	project = pom.SelectElement(keyProject)
	if nil == project {
		pom.CreateProcInst(keyXml, xmlDeclare)
		project = pom.CreateElement(keyProject)
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
		p.writeRepository(repositories, repo.id(), repo.Url, releaseFormat, keyRepository)
	}
}

func (p *stepPom) writeDistribution(project *etree.Element) {
	keyDistribution := "distributionManagement"
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	for _, repo := range p.Repositories {
		p.writeRepository(distribution, repo.id(), repo.Url, releaseFormat, keyRepository)
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
