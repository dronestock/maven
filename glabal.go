package main

import (
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/goexl/gfx"
)

func (p *plugin) global() (undo bool, err error) {
	if undo = 0 == len(p.Repositories); undo {
		return
	}

	filename := filepath.Join(os.Getenv(homeEnv), mavenHome, settingsFilename)
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

	// 配置全局
	settings := p.settings(doc)
	// 本地仓库
	p.writeLocalRepository(settings)
	// 组信息
	p.writeGroups(settings)
	// 镜像
	p.writeMirrors(settings)
	// 仓库
	p.writeServers(settings)
	// 配置
	p.writeProfiles(settings)

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}
