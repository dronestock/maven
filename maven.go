package main

import (
	"fmt"

	"github.com/dronestock/drone"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (p *plugin) mvn(args ...any) (err error) {
	fields := gox.Fields{
		field.String(`exe`, exe),
		field.String(`source`, p.Source),
		field.Any(`args`, args),
	}

	// 额外参数
	for key, value := range p.defines() {
		args = append(args, `--define`, fmt.Sprintf(`%s=%s`, key, value))
	}

	if err = p.Exec(exe, drone.Args(args...), drone.Dir(p.Source)); nil != err {
		p.Error(`执行出错`, fields.Connect(field.Any(`args`, args)).Connect(field.Error(err))...)
	} else if p.Verbose {
		p.Info(`执行成功`, fields...)
	}

	return
}
