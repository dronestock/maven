package main

import (
	"path/filepath"

	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

func (p *plugin) mvn(builder *args.Builder, repo *repository) (err error) {
	// 额外参数
	for key, value := range p.defines() {
		builder.Arg("define", gox.StringBuilder(key, equal, value).String())
	}
	// 强制使用UTF-8编码，避免乱码
	builder.Arg("define", gox.StringBuilder(fileEncoding, equal, utf8).String())
	// 全局配置文件
	if abs, ae := filepath.Abs(p.Filepath.Settings); nil == ae {
		builder.Arg("settings", abs)
	}
	// 指定模块文件
	if abs, ae := filepath.Abs(repo.filename(p.Source)); nil == ae && p.override() {
		builder.Arg("file", abs)
	}
	// 启用调试模式
	if p.Enabled(simaqian.LevelDebug) {
		builder.Flag("debug")
	}

	arguments := builder.Build()
	fields := gox.Fields[any]{
		field.New("exe", p.Binary.Maven),
		field.New("source", p.Source),
		field.New("args", arguments),
	}
	command := p.Command(p.Binary.Maven).Args(arguments).Dir(p.Source)
	p.Java.setHome(command)
	if _, err = command.Build().Exec(); nil != err {
		p.Error("Maven命令执行出错", fields.Add(field.Error(err))...)
	} else {
		p.Debug("Maven命令执行成功", fields...)
	}

	return
}
