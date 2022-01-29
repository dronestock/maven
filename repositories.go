package main

import (
	`fmt`

	`github.com/beevik/etree`
)

const (
	keyEnabled      = `enabled`
	keyUpdatePolicy = `updatePolicy`

	xmlCentral    = `central`
	xmlCentralUrl = `http://central`

	centralRepositoryFormat = `repository[url='%s']`
	centralPluginFormat     = `pluginRepository[url='%s']`
)

func (p *plugin) repositories(project *etree.Element) {
	repositories := project.SelectElement(keyRepositories)
	if nil == repositories {
		repositories = project.CreateElement(keyRepositories)
	}

	repository := repositories.FindElementPath(etree.MustCompilePath(fmt.Sprintf(centralRepositoryFormat, xmlCentralUrl)))
	if nil == repository {
		repository = repositories.CreateElement(keyRepository)
		repository.CreateElement(keyId).SetText(xmlCentral)
		repository.CreateElement(keyUrl).SetText(xmlCentralUrl)
		repository.CreateElement(keyReleases).CreateElement(keyEnabled).SetText(xmlTrue)
		repository.CreateElement(keySnapshots).CreateElement(keyEnabled).SetText(xmlTrue)
	}

	plugins := project.SelectElement(keyPluginRepositories)
	if nil == plugins {
		plugins = project.CreateElement(keyPluginRepositories)
	}

	_plugin := plugins.FindElementPath(etree.MustCompilePath(fmt.Sprintf(centralPluginFormat, xmlCentralUrl)))
	if nil == _plugin {
		_plugin = plugins.CreateElement(keyPluginRepository)
		_plugin.CreateElement(keyId).SetText(xmlCentral)
		_plugin.CreateElement(keyUrl).SetText(xmlCentralUrl)
		_plugin.CreateElement(keyReleases).CreateElement(keyEnabled).SetText(xmlTrue)

		snapshots := _plugin.CreateElement(keySnapshots)
		snapshots.CreateElement(keyEnabled).SetText(xmlTrue)
		snapshots.CreateElement(keyUpdatePolicy).SetText(xmlAlways)
	}
}
