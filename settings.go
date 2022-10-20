package main

import (
	"github.com/beevik/etree"
)

const keySettings = `settings`

func (p *plugin) settings(doc *etree.Document) (settings *etree.Element) {
	settings = doc.SelectElement(keySettings)
	if nil == settings {
		doc.CreateProcInst(keyXml, xmlDeclare)
		settings = doc.CreateElement(keySettings)
		settings.CreateAttr(keyXmlns, xmlSettingsXmlns)
		settings.CreateAttr(keyXsi, xmlSettingsXsi)
		settings.CreateAttr(keySchema, xmlSettingsSchema)
	}

	return
}
