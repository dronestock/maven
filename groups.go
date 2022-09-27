package main

import (
	"github.com/beevik/etree"
)

const (
	keySettingsGroups = `pluginGroups`
	keySettingsGroup  = `pluginGroup`

	xmlPlugins = `org.sonatype.plugins`
)

func (p *plugin) writeGroups(settings *etree.Element) {
	groups := settings.CreateElement(keySettingsGroups)
	if nil != groups {
		group := groups.CreateElement(keySettingsGroup)
		group.CreateText(xmlPlugins)
	}
}
