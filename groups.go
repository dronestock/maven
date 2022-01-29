package main

import (
	`github.com/beevik/etree`
)

const (
	keySettingsGroups = `pluginGroups`
	keySettingsGroup  = `pluginGroup`

	xmlPlugins = `org.sonatype.plugins`
)

func (p *plugin) groups(settings *etree.Element) {
	groups := settings.CreateElement(keySettingsGroups)
	if nil == groups {
		group := groups.CreateElement(keySettingsGroup)
		group.CreateText(xmlPlugins)
	}
}
