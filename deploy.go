package main

import (
	`strings`

	`github.com/dronestock/drone`
)

func (p *plugin) deploy() (undo bool, err error) {
	if undo = `` == strings.TrimSpace(p.Username) || `` == strings.TrimSpace(p.Password); undo {
		return
	}

	args := []interface{}{
		`deploy`,
	}
	// 打印更多日志
	if p.Verbose {
		args = append(args, `-X`)
	}

	// 执行命令
	err = p.Exec(exe, drone.Args(args...), drone.Dir(p.Folder))

	return
}
