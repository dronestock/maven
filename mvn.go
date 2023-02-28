package main

import (
	"fmt"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (p *plugin) mvn(args ...any) (err error) {
	fields := gox.Fields[any]{
		field.New("exe", exe),
		field.New("source", p.Source),
		field.New("args", args),
	}

	// 额外参数
	for key, value := range p.defines() {
		args = append(args, "--define", fmt.Sprintf("%s=%s", key, value))
	}

	if err = p.Command(exe).Args(args...).Dir(p.Source).Exec(); nil != err {
		p.Error("Maven命令执行出错", fields.Add(field.Error(err))...)
	} else if p.Verbose {
		p.Debug("Maven命令执行成功", fields...)
	}

	return
}
