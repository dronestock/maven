package main

import (
	`github.com/beevik/etree`
)

const (
	gpgPath = `plugin[artifactId='maven-gpg-plugin']`

	xmlPluginGpgArtifact = `maven-gpg-plugin`
	xmlPluginGpg         = `sign-artifacts`
)

func (p *plugin) sign(plugins *etree.Element) {
	if `` == p.Password {
		return
	}

	sign := plugins.FindElementPath(etree.MustCompilePath(gpgPath))
	if nil != sign {
		plugins.RemoveChildAt(sign.Index())
	}
	if !p.Sources {
		return
	}

	sign = plugins.CreateElement(keyPlugin)
	sign.CreateElement(keyGroupId).SetText(xmlMavenGroup)
	sign.CreateElement(keyArtifactId).SetText(xmlPluginGpgArtifact)
	sign.CreateElement(keyVersion).SetText(p.GpgPluginVersion)

	execution := sign.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginGpg)
	execution.CreateElement(keyPhase).SetText(xmlPluginVerify)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginSign)
}
