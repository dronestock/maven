package main

import (
	`github.com/beevik/etree`
)

const keyProject = `project`

func (p *plugin) project(doc *etree.Document) (project *etree.Element) {
	project = doc.SelectElement(keyProject)
	if nil == project {
		doc.CreateProcInst(keyXml, xmlDeclare)
		project = doc.CreateElement(keyProject)
		project.CreateAttr(keyXmlns, xmlProjectXmlns)
		project.CreateAttr(keyXsi, xmlProjectXsi)
		project.CreateAttr(keySchema, xmlProjectSchema)
	}

	return
}
