package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) do() (undo bool, err error) {
	args := make([]string, 0)

	// 清理
	if p.Clean {
		args = append(args, `clean`)
	}

	// 测试
	if p.Test {
		args = append(args, `test`)
	} else {
		args = append(args, `-Dmaven.skip.test=true`)
	}

	// 打包
	args = append(args, `package`)

	// 发布
	if `` != p.Username && `` != p.Password {
		args = append(args, `deploy`)
	}

	// 执行命令
	err = p.Exec(exe, drone.Args(args...), drone.Dir(p.Folder))

	return
}
