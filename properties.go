package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) properties(project *etree.Element) {
	properties := project.SelectElement(keyProperties)
	if nil == properties {
		properties = project.CreateElement(keyProperties)
	}

	for key, value := range p._properties() {
		properties.CreateElement(key).SetText(value)
	}
}
