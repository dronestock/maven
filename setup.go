package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) setup(project *etree.Element) {
	group := project.SelectElement(keyGroup)
	if nil == group {
		group = project.CreateElement(keyGroup)
	}
	if `` != p.Group {
		group.SetText(p.Group)
	}

	artifact := project.SelectElement(keyArtifact)
	if nil == artifact {
		artifact = project.CreateElement(keyArtifact)
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
