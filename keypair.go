package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) keypair() (undo bool, err error) {
	args := []string{
		`--batch`,
		`--passphrase`,
		p.Password,
		`--quick-gen-key`,
		p.Username,
		`default`,
		`default`,
	}
	// 执行命令
	err = p.Exec(gpgExe, drone.Args(args...), drone.Dir(p.Folder))
	// gpg --keyserver hkp://keyserver.ubuntu.com --send-keys
	// $(gpg --list-signatures --with-colons | grep 'sig' | grep 'the Name-Real of my key' | head -n 1 | cut -d':' -f5)

	return
}
