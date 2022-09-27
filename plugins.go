package main

import (
	"github.com/beevik/etree"
)

const (
	keyBuild   = `build`
	keyPlugins = `plugins`
	keyPlugin  = `plugin`

	keyExecutions    = `executions`
	keyExecution     = `execution`
	keyPhase         = `phase`
	keyGoals         = `goals`
	keyGoal          = `goal`
	keyConfiguration = `configuration`

	xmlPluginVerify    = `verify`
	xmlPluginJar       = `jar`
	xmlPluginTestJar   = `test-jar`
	xmlPluginJarNoFork = `jar-no-fork`
	xmlPluginSign      = `sign`
)

func (p *plugin) writePlugins(project *etree.Element) {
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
