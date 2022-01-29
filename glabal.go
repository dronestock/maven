package main

import (
	`os`
	`path/filepath`

	`github.com/beevik/etree`
	`github.com/storezhang/gfx`
)

func (p *plugin) global() (undo bool, err error) {
	if undo = `` == p.Username || `` == p.Password; undo {
		return
	}

	filename := filepath.Join(os.Getenv(homeEnv), mavenDir, settingsFilename)
	if !gfx.Exist(filename) {
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

	// 配置全局
	settings := p.settings(doc)
	// 组信息
	p.groups(settings)
	// 镜像
	// p.mirrors(settings)
	// 仓库
	p.servers(settings)
	// 配置
	p.profiles(settings)

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}
