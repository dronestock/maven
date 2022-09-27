package main

import (
	"fmt"

	"github.com/beevik/etree"
)

const (
	repositoryFormat         = `repository[id='%s']`
	snapshotRepositoryFormat = `snapshotRepository[id='%s']`

	keyDistribution = `distributionManagement`
)

func (p *plugin) writeDistribution(project *etree.Element) {
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	for _, _repository := range p.Repositories {
		// 写入正式仓库
		p.writeReleaseRepository(distribution, _repository)
		// 写入快照仓库
		p.writeSnapshotRepository(distribution, _repository)
	}

}

func (p *plugin) writeReleaseRepository(distribution *etree.Element, repository *repository) {
	id := repository.releaseId()
	releasePath := etree.MustCompilePath(fmt.Sprintf(repositoryFormat, id))
	_repository := distribution.FindElementPath(releasePath)
	if nil != _repository {
		distribution.RemoveChildAt(_repository.Index())
	}
	_repository = distribution.CreateElement(keyRepository)
	if _repository.Text() != id {
		_repository.CreateElement(keyId).SetText(id)
		_repository.CreateElement(keyName).SetText(id)
		_repository.CreateElement(keyUrl).SetText(repository.Release)
	}
}

func (p *plugin) writeSnapshotRepository(distribution *etree.Element, repository *repository) {
	id := repository.snapshotId()
	snapshotPath := etree.MustCompilePath(fmt.Sprintf(snapshotRepositoryFormat, id))
	snapshot := distribution.FindElementPath(snapshotPath)
	if nil != snapshot {
		distribution.RemoveChildAt(snapshot.Index())
	}
	snapshot = distribution.CreateElement(keySnapshotRepository)
	if snapshot.Text() != id {
		snapshot.CreateElement(keyId).SetText(id)
		snapshot.CreateElement(keyName).SetText(id)
		snapshot.CreateElement(keyUrl).SetText(repository.Snapshot)
	}
}
