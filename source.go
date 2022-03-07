package main

import (
	`github.com/beevik/etree`
)

const (
	sourcePath = `plugin[artifactId='maven-source-plugin']`

	xmlPluginSourceArtifact = `maven-source-plugin`
	xmlPluginSource         = `attach-source`
)

func (p *plugin) source(plugins *etree.Element) {
	sources := plugins.FindElementPath(etree.MustCompilePath(sourcePath))
	if nil != sources {
		plugins.RemoveChildAt(sources.Index())
	}
	if !p.Source {
		return
	}

	sources = plugins.CreateElement(keyPlugin)
	sources.CreateElement(keyGroupId).SetText(xmlMavenGroup)
	sources.CreateElement(keyArtifactId).SetText(xmlPluginSourceArtifact)
	sources.CreateElement(keyVersion).SetText(p.SourcePluginVersion)

	execution := sources.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginSource)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginJarNoFork)
}
