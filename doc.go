package main

import (
	"github.com/beevik/etree"
)

const (
	docPath = `plugin[artifactId='maven-javadoc-plugin']`

	xmlPluginDocArtifact = `maven-javadoc-plugin`
	xmlPluginDoc         = `attach-javadocs`
)

func (p *plugin) writeDoc(plugins *etree.Element) {
	docs := plugins.FindElementPath(etree.MustCompilePath(docPath))
	if nil != docs {
		plugins.RemoveChildAt(docs.Index())
	}
	if !p.Docs {
		return
	}

	docs = plugins.CreateElement(keyPlugin)
	docs.CreateElement(keyGroupId).SetText(xmlMavenGroup)
	docs.CreateElement(keyArtifactId).SetText(xmlPluginDocArtifact)
	docs.CreateElement(keyVersion).SetText(p.DocPluginVersion)

	execution := docs.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginDoc)
	execution.CreateElement(keyGoals).CreateElement(keyGoal).SetText(xmlPluginJar)
}
