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
	server := servers.FindElementPath(etree.MustCompilePath(fmt.Sprintf(serverPathFormat, p.repositoryId())))
	if nil != server {
		servers.RemoveChildAt(server.Index())
	}

	server = servers.CreateElement(keyServer)
	server.CreateElement(keyId).SetText(p.repositoryId())
	server.CreateElement(keyUsername).SetText(p.Username)
	server.CreateElement(keyPassword).SetText(p.Password)
}
