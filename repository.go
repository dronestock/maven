package main

import (
	"net/url"

	"github.com/goexl/gox"
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
}

func (r *repository) snapshotId() string {
	return r._id(r.Snapshot)
}

func (r *repository) releaseId() string {
	return r._id(r.Release)
}

func (r *repository) _id(link string) (id string) {
	if uri, err := url.Parse(link); nil != err {
		id = gox.RandString(randLength)
	} else {
		id = uri.Host
	}

	return
}
