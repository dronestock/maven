package main

import (
	`fmt`
	`os`
	`path/filepath`

	`github.com/dronestock/drone`
	`github.com/storezhang/gfx`
)

const listKeyFormat = `$(gpg --list-signatures --with-colons | grep 'sig' | grep '%s' | head -n 1 | cut -d':' -f5)`

func (p *plugin) keypair() (undo bool, err error) {
	// 删除原来的密钥目录
	if err = gfx.Delete(filepath.Join(os.Getenv(homeEnv), gpgHome)); nil != err {
		return
	}

	// 生成密钥
	makeArgs := []string{
		`--batch`,
		`--passphrase`,
		p.Password,
		`--quick-gen-key`,
		p.Username,
		`default`,
		`default`,
	}
	if err = p.Exec(gpgExe, drone.Args(makeArgs...), drone.Dir(p.Folder)); nil != err {
		return
	}

	// 上传到服务器
	sendArgs := []string{
		`--keyserver`,
		p.GpgServer,
		`--send-keys`,
		fmt.Sprintf(listKeyFormat, p.Username),
	}
	// TODO 这儿有问题，后续解决
	_ = p.Exec(gpgExe, drone.Args(sendArgs...), drone.Dir(p.Folder))

	return
}
