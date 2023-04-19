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
	switch j.Version {
	case javaVersionLts:
		command.Environment().Kv(javaHome, j.Lts)
	case javaVersionLatest:
		command.Environment().Kv(javaHome, j.Latest)
	}

	return
}
