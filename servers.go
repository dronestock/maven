package main

import (
	"fmt"

	"github.com/beevik/etree"
)

const (
	serverPathFormat = `server[id='%s']`

	keyServers = `servers`
	keyServer  = `server`

	keyUsername = `username`
	keyPassword = `password`
)

func (p *plugin) writeServers(settings *etree.Element) {
	servers := settings.CreateElement(keyServers)

	for _, _repository := range p.Repositories {
		// 写入正式服务器
		p.writeReleaseServer(servers, _repository)
		// 写入快照服务器
		p.writeSnapshotServer(servers, _repository)
	}
}

func (p *plugin) writeReleaseServer(servers *etree.Element, repository *repository) {
	path := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, repository.releaseId()))
	release := servers.FindElementPath(path)
	if nil != release {
		servers.RemoveChildAt(release.Index())
	}
	release = servers.CreateElement(keyServer)
	release.CreateElement(keyId).SetText(repository.releaseId())
	release.CreateElement(keyUsername).SetText(repository.Username)
	release.CreateElement(keyPassword).SetText(repository.Password)
}

func (p *plugin) writeSnapshotServer(servers *etree.Element, repository *repository) {
	path := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, repository.snapshotId()))
	snapshot := servers.FindElementPath(path)
	if nil != snapshot {
		servers.RemoveChildAt(snapshot.Index())
	}
	snapshot = servers.CreateElement(keyServer)
	snapshot.CreateElement(keyId).SetText(repository.snapshotId())
	snapshot.CreateElement(keyUsername).SetText(repository.Username)
	snapshot.CreateElement(keyPassword).SetText(repository.Password)
}
