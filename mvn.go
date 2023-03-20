package main

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (p *plugin) mvn(builder *args.Builder) (err error) {
	// 额外参数
	for key, value := range p.defines() {
		builder.Args("define", gox.StringBuilder(key, equal, value).String())
	}

	arguments := builder.Build()
	fields := gox.Fields[any]{
		field.New("exe", exe),
		field.New("source", p.Source),
		field.New("args", arguments),
	}
	if _, err = p.Command(exe).Args(arguments).Dir(p.Source).Build().Exec(); nil != err {
		p.Error("Maven命令执行出错", fields.Add(field.Error(err))...)
	} else {
		p.Debug("Maven命令执行成功", fields...)
	}

	return
}
