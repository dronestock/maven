package main

import (
	`os`
	`path/filepath`

	`github.com/dronestock/drone`
)

func (p *plugin) keypair() (undo bool, err error) {
	// 删除原来的密钥目录
	if err = gfx.Delete(filepath.Join(os.Getenv(homeEnv), gpgHome)); nil != err {
		return
	}

	args := []interface{}{
		`--batch`,
		`--passphrase`,
		p.Password,
		`--quick-gen-key`,
		p.Username,
		`default`,
		`default`,
		p.Gpg.Expire,
	}
	err = p.Exec(gpgExe, drone.Args(args...), drone.Dir(p.Source))

	return
}
