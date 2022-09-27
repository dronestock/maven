package main

import (
	"github.com/beevik/etree"
)

const (
	jarPath = `plugin[artifactId='maven-jar-plugin']`

	xmlPluginJarArtifact = `maven-jar-plugin`
	xmlPluginDefaultJar  = `default-jar`
	xmlPluginPackage     = `package`
)

func (p *plugin) writeJar(plugins *etree.Element) {
	jar := plugins.FindElementPath(etree.MustCompilePath(jarPath))
	if nil != jar {
		plugins.RemoveChildAt(jar.Index())
	}
	if !p.Sources {
		return
	}

	jar = plugins.CreateElement(keyPlugin)
	jar.CreateElement(keyGroupId).SetText(xmlMavenGroup)
	jar.CreateElement(keyArtifactId).SetText(xmlPluginJarArtifact)
	jar.CreateElement(keyVersion).SetText(p.JarPluginVersion)

	execution := jar.CreateElement(keyExecutions).CreateElement(keyExecution)
	execution.CreateElement(keyId).SetText(xmlPluginDefaultJar)
	execution.CreateElement(keyPhase).SetText(xmlPluginPackage)

	goals := execution.CreateElement(keyGoals)
	goals.CreateElement(keyGoal).SetText(xmlPluginJar)
	goals.CreateElement(keyGoal).SetText(xmlPluginTestJar)
}
