package main

type gpg struct {
	// 服务器
	Server string `default:"hkp://keyserver.ubuntu.com"`
	// 过期时间
	Expire string `default:"7d"`
}
