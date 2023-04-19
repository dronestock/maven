package main

import (
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (p *plugin) mvn(builder *args.Builder) (err error) {
	// 额外参数
	for key, value := range p.defines() {
		builder.Arg("define", gox.StringBuilder(key, equal, value).String())
	}
	// 全局配置文件
	if abs, ae := filepath.Abs(p.Filepath.Settings); nil == ae {
		builder.Arg("settings", abs)
	}
	// 指定模块文件
	if abs, ae := filepath.Abs(p.filename); nil == ae && p.override() {
		builder.Arg("file", abs)
	}

	arguments := builder.Build()
	fields := gox.Fields[any]{
		field.New("exe", p.Binary.Maven),
		field.New("source", p.Source),
		field.New("args", arguments),
	}
	if _, err = p.Command(p.Binary.Maven).Args(arguments).Dir(p.Source).Build().Exec(); nil != err {
		p.Error("Maven命令执行出错", fields.Add(field.Error(err))...)
	} else {
		p.Debug("Maven命令执行成功", fields...)
	}

	return
}
