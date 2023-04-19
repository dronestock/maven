package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/beevik/etree"
	"github.com/goexl/gox"
	"github.com/goexl/gox/rand"
)

const fullRepositoryFormat = "/project/repositories/repository[url='%s']"

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
	// 是否需要部署
	Deploy *bool `default:"true" json:"deploy,omitempty"`
}

func (r *repository) snapshotId(doc *etree.Document) (id string) {
	if id = r.savedId(doc, r.Snapshot); "" == id {
		id = r.id(r.Snapshot, snapshot)
	}

	return
}

func (r *repository) releaseId(doc *etree.Document) (id string) {
	if id = r.savedId(doc, r.Release); "" == id {
		id = r.id(r.Release, release)
	}

	return
}

func (r *repository) private() bool {
	return !strings.HasPrefix(r.Snapshot, mavenRepositoryHost) || r.Private
}

func (r *repository) id(link string, suffix string) (id string) {
	if uri, err := url.Parse(link); nil != err {
		id = rand.New().String().Length(randLength).Build().Generate()
	} else {
		id = gox.StringBuilder(uri.Host, dot, suffix).String()
	}

	return
}

func (r *repository) savedId(doc *etree.Document, url string) (id string) {
	path := fmt.Sprintf(fullRepositoryFormat, url)
	_repository := doc.FindElementPath(etree.MustCompilePath(path))
	if nil != _repository {
		idElement := _repository.FindElement(keyId)
		id = gox.If(nil != idElement, idElement.Text())
	}

	return
}
