package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) sources(plugins *etree.Element) {
	sources := plugins.FindElementPath(etree.MustCompilePath(sourcesPath))
	if nil != sources {
		plugins.RemoveChildAt(sources.Index())
	}
	if !p.Sources {
		return
	}

	sources = plugins.CreateElement(keyPlugin)
	sources.CreateElement(keyArtifact).SetText(xmlPluginSourceArtifact)
	execution := sources.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginSources)
	execution.CreateElement(keyPhase).SetText(xmlPluginPackage)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginJarNoFork)
}
