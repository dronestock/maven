package main

import (
	`fmt`

	`github.com/beevik/etree`
)

const repositoryFormat = `repository[id='%s']`

func (p *plugin) distribution(project *etree.Element) {
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	repository := distribution.FindElementPath(etree.MustCompilePath(fmt.Sprintf(repositoryFormat, p.repositoryId())))
	if nil != repository {
		distribution.RemoveChildAt(repository.Index())
	}

	repository = distribution.CreateElement(keyRepository)
	if repository.Text() != p.repositoryId() {
		repository.CreateElement(keyId).SetText(p.repositoryId())
		repository.CreateElement(keyName).SetText(p.repositoryId())
		repository.CreateElement(keyUrl).SetText(p.Repository)
	}
}
