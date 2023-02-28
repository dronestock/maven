package main

import (
	"net/url"
	"strings"

	"github.com/goexl/gox/rand"
)

type repository struct {
	// 正式仓库
	// nolint:lll
	Release string `default:"https://s01.oss.sonatype.org/service/local/staging/deploy/maven2" json:"release" validate:"required"`
	// 快照仓库
	// nolint:lll
	Snapshot string `default:"https://s01.oss.sonatype.org/content/writeRepositories/snapshots" json:"snapshot" validate:"required"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 是否为私服，不对外开放
	Private bool `default:"true" json:"private"`
}

func (r *repository) snapshotId() string {
	return r.id(r.Snapshot)
}

func (r *repository) releaseId() string {
	return r.id(r.Release)
}

func (r *repository) private() bool {
	return !strings.HasPrefix(r.Snapshot, mavenRepositoryHost) || r.Private
}

func (r *repository) id(link string) (id string) {
	if uri, err := url.Parse(link); nil != err {
		id = rand.New().String().Length(randLength).Generate()
	} else {
		id = uri.Host
	}

	return
}
