package main

func (p *plugin) keypair() (undo bool, err error) {
	// 删除原来的密钥目录
	/*if err = gfx.Delete(filepath.Join(os.Getenv(homeEnv), gpgHome)); nil != err {
		return
	}

	for _, _repository := range p.Repositories {
		args := []interface{}{
			`--batch`,
			`--passphrase`,
			_repository.Password,
			`--quick-gen-key`,
			_repository.Username,
			`default`,
			`default`,
			p.Gpg.Expire,
		}
		err = p.Exec(gpgExe, drone.Args(args...), drone.Dir(p.Source))
		if nil != err {
			return
		}
	}*/

	return
}
