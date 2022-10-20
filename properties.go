package main

import (
	"github.com/beevik/etree"
)

func (p *plugin) writeProperties(project *etree.Element) {
	properties := project.SelectElement(keyProperties)
	if nil == properties {
		properties = project.CreateElement(keyProperties)
	}

	for key, value := range p.properties() {
		property := properties.SelectElement(key)
		if nil == property {
			property = properties.CreateElement(key)
		}
		property.SetText(value)
	}
}
