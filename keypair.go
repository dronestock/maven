package main

import (
	`os`
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gfx`
)

func (p *plugin) keypair() (undo bool, err error) {
	// 删除原来的密钥目录
	if err = gfx.Delete(filepath.Join(os.Getenv(homeEnv), gpgHome)); nil != err {
		return
	}

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
