package main

import (
	`fmt`

	`github.com/beevik/etree`
)

const (
	serverPathFormat = `server[id='%s']`

	keyServers = `servers`
	keyServer  = `server`

	keyUsername = `username`
	keyPassword = `password`
)

func (p *plugin) servers(settings *etree.Element) () {
	servers := settings.CreateElement(keyServers)

	// 写入正式服务器
	releasePath := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, p.repositoryId(p.Repository.Release)))
	release := servers.FindElementPath(releasePath)
	if nil != release {
		servers.RemoveChildAt(release.Index())
	}
	release = servers.CreateElement(keyServer)
	release.CreateElement(keyId).SetText(p.repositoryId(p.Repository.Release))
	release.CreateElement(keyUsername).SetText(p.Username)
	release.CreateElement(keyPassword).SetText(p.Password)

	// 写入快照服务器
	snapshotPath := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, p.repositoryId(p.Repository.Snapshot)))
	snapshot := servers.FindElementPath(snapshotPath)
	if nil != snapshot {
		servers.RemoveChildAt(snapshot.Index())
	}
	snapshot = servers.CreateElement(keyServer)
	snapshot.CreateElement(keyId).SetText(p.repositoryId(p.Repository.Snapshot))
	snapshot.CreateElement(keyUsername).SetText(p.Username)
	snapshot.CreateElement(keyPassword).SetText(p.Password)
}
