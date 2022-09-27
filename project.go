package main

import (
	"github.com/beevik/etree"
)

const keyProject = `project`

func (p *plugin) writeProject(doc *etree.Document) (project *etree.Element) {
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
	if `` != p.Group {
		group.SetText(p.Group)
	}

	artifact := project.SelectElement(keyArtifactId)
	if nil == artifact {
		artifact = project.CreateElement(keyArtifactId)
	}
	if `` != p.Artifact {
		artifact.SetText(p.Artifact)
	}

	version := project.SelectElement(keyVersion)
	if nil == version {
		version = project.CreateElement(keyVersion)
	}
	if `` != p.Version {
		version.SetText(p.Version)
	}

	packaging := project.SelectElement(keyPackaging)
	if nil == packaging {
		packaging = project.CreateElement(keyPackaging)
	}
	if `` != p.Packaging {
		packaging.SetText(p.Packaging)
	}

	return
}
