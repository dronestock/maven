package main

import (
	"fmt"

	"github.com/dronestock/drone"
)

func (p *plugin) pkg() (undo bool, err error) {
	args := make([]interface{}, 0)

	// 清理
	if p.Clean {
		args = append(args, `clean`)
	}

	// 测试
	if p.Test {
		args = append(args, `test`)
	}

	// 打包
	args = append(args, `package`)

	// 测试参数
	if !p.Test {
		args = append(args, `--define`, `maven.skip.test=true`)
	}
	// 额外参数
	for key, value := range p.Defines {
		args = append(args, `--define`, fmt.Sprintf(`%s=%s`, key, value))
	}

	// 打印更多日志
	if p.Verbose {
		args = append(args, `-X`)
	}

	// 执行命令
	err = p.Exec(exe, drone.Args(args...), drone.Dir(p.Source))

	return
}
