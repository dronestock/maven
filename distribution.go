package main

import (
	`github.com/beevik/etree`
)

func (p *plugin) distribution(project *etree.Element) {
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}
	repository := distribution.CreateElement(keyRepository)
	repository.CreateElement(keyId).SetText(repositoryId)
	repository.CreateElement(keyName).SetText(repositoryId)
	repository.CreateElement(keyUrl).SetText(p.Repository)
}
