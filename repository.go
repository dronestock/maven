package main

import (
	"net/url"
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/rand"
)

type repository struct {
	// 地址
	Url string `default:"https://central.sonatype.com" json:"url" validate:"required"`
	// 用户名
	Username string `json:"username" validate:"required"`
	// 密码
	Password string `json:"password" validate:"required"`
	// 是否为私服，不对外开放
	Private bool `json:"private"`
	// 是否需要部署
	Deploy *bool `default:"true" json:"deploy,omitempty"`
}

func (r *repository) id() (id string) {
	if uri, err := url.Parse(r.Url); nil != err {
		id = rand.New().String().Length(randLength).Build().Generate()
	} else {
		id = gox.StringBuilder(uri.Host).String()
	}

	return
}

func (r *repository) filename(source string) (filename string) {
	pom := gox.StringBuilder(r.id(), dot, pomFilename).String()
	filename = filepath.Join(source, pom)

	return
}
