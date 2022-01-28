package main

import (
	`path/filepath`

	`github.com/beevik/etree`
	`github.com/storezhang/gfx`
)

func (p *plugin) pom() (undo bool, err error) {
	pom := filepath.Join(p.Folder, pomFilename)
	if err = gfx.Delete(pom); nil != err {
		return
	}

	dir := filepath.Dir(pom)
	if err = gfx.Create(dir, gfx.Dir()); nil != err {
		return
	}
	if err = gfx.Create(pom, gfx.File()); nil != err {
		return
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromFile(pom); nil != err {
		return
	}

	// 设置项目
	project := p.project(doc)
	// 设置坐标
	p.setup(project)
	// 设置项目属性
	p.properties(project)
	// 设置发布仓库
	p.distribution(project)
	// 设置发布插件
	p.plugins(project)

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(pom)

	return
}
