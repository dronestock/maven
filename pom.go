package main

import (
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/goexl/gfx"
)

func (p *plugin) pom() (undo bool, err error) {
	filename := filepath.Join(p.Source, pomFilename)
	if _, exists := gfx.Exists(filename); !exists {
		if err = gfx.Create(filepath.Dir(filename), gfx.Dir()); nil != err {
			return
		}
		if err = gfx.Create(filename, gfx.File()); nil != err {
			return
		}
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromFile(filename); nil != err {
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
	err = doc.WriteToFile(filename)

	return
}
