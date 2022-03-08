package main

import (
	`fmt`

	`github.com/beevik/etree`
)

const (
	repositoryFormat         = `repository[id='%s']`
	snapshotRepositoryFormat = `snapshotRepository[id='%s']`

	keyDistribution = `distributionManagement`
)

func (p *plugin) distribution(project *etree.Element) {
	distribution := project.SelectElement(keyDistribution)
	if nil == distribution {
		distribution = project.CreateElement(keyDistribution)
	}

	// 写入正式仓库
	releasePath := etree.MustCompilePath(fmt.Sprintf(repositoryFormat, p.repositoryId(p.Repository.Snapshot)))
	_repository := distribution.FindElementPath(releasePath)
	if nil != _repository {
		distribution.RemoveChildAt(_repository.Index())
	}
	_repository = distribution.CreateElement(keyRepository)
	if _repository.Text() != p.repositoryId(p.Repository.Release) {
		_repository.CreateElement(keyId).SetText(p.repositoryId(p.Repository.Release))
		_repository.CreateElement(keyName).SetText(p.repositoryId(p.Repository.Release))
		_repository.CreateElement(keyUrl).SetText(p.Repository.Release)
	}

	// 写入快照仓库
	snapshotPath := etree.MustCompilePath(fmt.Sprintf(snapshotRepositoryFormat, p.repositoryId(p.Repository.Snapshot)))
	snapshot := distribution.FindElementPath(snapshotPath)
	if nil != snapshot {
		distribution.RemoveChildAt(snapshot.Index())
	}
	snapshot = distribution.CreateElement(keySnapshotRepository)
	if snapshot.Text() != p.repositoryId(p.Repository.Snapshot) {
		snapshot.CreateElement(keyId).SetText(p.repositoryId(p.Repository.Snapshot))
		snapshot.CreateElement(keyName).SetText(p.repositoryId(p.Repository.Snapshot))
		snapshot.CreateElement(keyUrl).SetText(p.Repository.Snapshot)
	}
}
