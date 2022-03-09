package main

import (
	`github.com/beevik/etree`
)

const (
	nexusPath = `plugin[artifactId='nexus-staging-maven-plugin']`

	keyServerId    = `serverId`
	keyNexusUrl    = `nexusUrl`
	keyAutoRelease = `autoReleaseAfterClose`

	xmlNexusGroup          = `org.sonatype.plugins`
	xmlPluginNexusArtifact = `nexus-staging-maven-plugin`

	xmlNexusUrl = `https://oss.sonatype.org/`
)

func (p *plugin) nexus(plugins *etree.Element) {
	nexus := plugins.FindElementPath(etree.MustCompilePath(nexusPath))
	if nil != nexus {
		plugins.RemoveChildAt(nexus.Index())
	}
	if !p.Sources {
		return
	}

	nexus = plugins.CreateElement(keyPlugin)
	nexus.CreateElement(keyGroupId).SetText(xmlNexusGroup)
	nexus.CreateElement(keyArtifactId).SetText(xmlPluginNexusArtifact)
	nexus.CreateElement(keyVersion).SetText(p.NexusPluginVersion)

	configuration := nexus.CreateElement(keyConfiguration)
	configuration.CreateElement(keyServerId).SetText(keyRepository)
	configuration.CreateElement(keyNexusUrl).SetText(xmlNexusUrl)
	configuration.CreateElement(keyAutoRelease).SetText(xmlTrue)
}
