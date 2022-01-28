package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) plugins(project *etree.Element) {
	build := project.SelectElement(keyBuild)
	if nil == build {
		build = project.CreateElement(keyBuild)
	}
	plugins := build.SelectElement(keyPlugins)
	if nil == plugins {
		plugins = build.CreateElement(keyPlugins)
	}

	// 设置源码发布
	p.sources(plugins)
	// 设置文档发布
	p.docs(plugins)
}
