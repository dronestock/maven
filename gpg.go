package main

import (
	`github.com/beevik/etree`
)

const (
	gpgPath = `plugin[artifactId='maven-gpg-plugin']`

	xmlPluginGpgArtifact = `maven-gpg-plugin`
	xmlPluginGpg         = `sign-artifacts`
)

func (p *plugin) gpg(plugins *etree.Element) {
	gpg := plugins.FindElementPath(etree.MustCompilePath(gpgPath))
	if nil != gpg {
		plugins.RemoveChildAt(gpg.Index())
	}
	if !p.Sources {
		return
	}

	gpg = plugins.CreateElement(keyPlugin)
	gpg.CreateElement(keyGroupId).SetText(xmlMavenGroup)
	gpg.CreateElement(keyArtifactId).SetText(xmlPluginGpgArtifact)
	gpg.CreateElement(keyVersion).SetText(p.GpgPluginVersion)

	execution := gpg.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginGpg)
	execution.CreateElement(keyPhase).SetText(xmlPluginVerify)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginSign)
}
