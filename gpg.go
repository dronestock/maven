package main

type gpg struct {
	// 服务器
	Server string `default:"keys.openpgp.org" json:"server"`
	// 过期时间
	Expire string `default:"7d" json:"expire"`
	// 密码
	Passphrase string `json:"passphrase"`
}
