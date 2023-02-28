package main

import (
	"github.com/beevik/etree"
)

const (
	keyBuild           = "build"
	keyPlugins         = "plugins"
	keyPlugin          = "plugin"
	keyExecutions      = "executions"
	keyExecution       = "execution"
	keyPhase           = "phase"
	keyGoals           = "goals"
	keyGoal            = "goal"
	keyConfiguration   = "configuration"
	xmlPluginVerify    = "verify"
	xmlPluginJar       = "jar"
	xmlPluginTestJar   = "test-jar"
	xmlPluginJarNoFork = "jar-no-fork"
	xmlPluginSign      = "sign"

	jarPath              = "plugin[artifactId='maven-jar-plugin']"
	xmlPluginJarArtifact = "maven-jar-plugin"
	xmlPluginDefaultJar  = "default-jar"
	xmlPluginPackage     = "package"

	sourcePath              = "plugin[artifactId='maven-source-plugin']"
	xmlPluginSourceArtifact = "maven-source-plugin"
	xmlPluginSource         = "attach-source"

	docPath              = "plugin[artifactId='maven-javadoc-plugin']"
	xmlPluginDocArtifact = "maven-javadoc-plugin"
	xmlPluginDoc         = "attach-javadocs"

	gpgPath              = "plugin[artifactId='maven-gpg-plugin']"
	xmlPluginGpgArtifact = "maven-gpg-plugin"
	xmlPluginGpg         = "sign-artifacts"

	nexusPath              = "plugin[artifactId='nexus-staging-maven-plugin']"
	keyServerId            = "serverId"
	keyNexusUrl            = "nexusUrl"
	keyAutoRelease         = "autoReleaseAfterClose"
	xmlNexusGroup          = "org.sonatype.plugins"
	xmlPluginNexusArtifact = "nexus-staging-maven-plugin"
	xmlNexusUrl            = "https://oss.sonatype.org/"
)

func (p *stepPom) writePlugins(project *etree.Element) {
	build := project.SelectElement(keyBuild)
	if nil == build {
		build = project.CreateElement(keyBuild)
	}
	plugins := build.SelectElement(keyPlugins)
	if nil == plugins {
		plugins = build.CreateElement(keyPlugins)
	}

	// 设置打包
	p.writeJar(plugins)
	// 设置源码发布
	p.writeSource(plugins)
	// 设置文档发布
	p.writeDoc(plugins)
	// 设置构件签名
	p.writeSign(plugins)
	// 设置发布
	p.writeNexus(plugins)
}

func (p *stepPom) writeJar(plugins *etree.Element) {
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

func (p *stepPom) writeSource(plugins *etree.Element) {
	sources := plugins.FindElementPath(etree.MustCompilePath(sourcePath))
	if nil != sources {
		plugins.RemoveChildAt(sources.Index())
	}
	if !p.Sources {
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

func (p *stepPom) writeDoc(plugins *etree.Element) {
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

func (p *stepPom) writeSign(plugins *etree.Element) {
	if 0 == len(p.Repositories) {
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

func (p *stepPom) writeNexus(plugins *etree.Element) {
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
