package main

import (
	`fmt`
	`os`
	`path/filepath`

	`github.com/beevik/etree`
	`github.com/storezhang/gfx`
)

func (p *plugin) settings() (undo bool, err error) {
	if undo = `` == p.Username || `` == p.Password; undo {
		return
	}

	filename := filepath.Join(os.Getenv(homeEnv), mavenDir, settingsFilename)
	if err = gfx.Delete(filename); nil != err {
		return
	}

	dir := filepath.Dir(filename)
	if err = gfx.Create(dir, gfx.Dir()); nil != err {
		return
	}
	if err = gfx.Create(filename, gfx.File()); nil != err {
		return
	}

	doc := etree.NewDocument()
	doc.CreateProcInst(keyXml, xmlDeclare)

	// 命名空间
	settings := doc.CreateElement(keySettings)
	settings.CreateAttr(keyXmlns, xmlSettingsXmlns)
	settings.CreateAttr(keyXsi, xmlSettingsXsi)
	settings.CreateAttr(keySchema, xmlSettingsSchema)

	// 组信息
	groups := settings.CreateElement(keySettingsGroups)
	group := groups.CreateElement(keySettingsGroup)
	group.CreateText(`org.sonatype.plugins`)

	// 镜像
	if 0 < len(p.Mirrors) {
		mirrors := settings.CreateElement(keyMirrors)
		count := 1
		for _, url := range p.Mirrors {
			mirror := mirrors.CreateElement(keyMirror)
			mirror.CreateElement(keyId).CreateText(fmt.Sprintf(toIntFormat, count))
			mirror.CreateElement(keyMirrorOf).CreateText(`*`)
			mirror.CreateElement(keyName).CreateText(fmt.Sprintf(toIntFormat, count))
			mirror.CreateElement(keyUrl).CreateText(url)

			count++
		}
	}

	// 仓库
	servers := settings.CreateElement(keyServers)
	repository := servers.CreateElement(keyServer)
	repository.CreateElement(keyId).SetText(repositoryId)
	repository.CreateElement(keyUsername).SetText(p.Username)
	repository.CreateElement(keyPassword).SetText(p.Password)

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}
