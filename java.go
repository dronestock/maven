package main

import (
	"github.com/goexl/gex"
)

type java struct {
	// 版本
	Version javaVersion `default:"${JAVA_VERSION=lts}" json:"version,omitempty" validate:"oneof=lts latest"`
	// 长期版本路径
	Lts string `default:"${JAVA_LTS}" json:"lts,omitempty"`
	// 最新版本路径
	Latest string `default:"${JAVA_LATEST}" json:"latest,omitempty"`
}

func (j *java) setHome(command *gex.Builder) {
	home := ""
	if javaVersionLts == j.Version {
		home = j.Lts
	} else if javaVersionLatest == j.Version {
		home = j.Latest
	}
	if "" != home {
		command.Environment().Kv(javaHome, home).Build()
	}
}
