package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) docs(plugins *etree.Element) {
	docs := plugins.FindElementPath(etree.MustCompilePath(docsPath))
	if nil != docs {
		plugins.RemoveChildAt(docs.Index())
	}
	if !p.Docs {
		return
	}

	docs = plugins.CreateElement(keyPlugin)
	docs.CreateElement(keyArtifact).SetText(xmlPluginDocArtifact)
	execution := docs.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginDocs)
	execution.CreateElement(keyPhase).SetText(xmlPluginPackage)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginJar)
}
