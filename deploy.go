package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) deploy() (undo bool, err error) {
	if undo = 0 == len(p.Repositories); undo {
		return
	}

	args := []interface{}{
		`deploy`,
		`--define`, `maven.wagon.http.ssl.insecure=true`,
		`--define`, `maven.wagon.http.ssl.allowall=true`,
	}
	// 打印更多日志
	if p.Verbose {
		args = append(args, `-X`)
	}

	// 执行命令
	err = p.Exec(exe, drone.Args(args...), drone.Dir(p.Source))

	return
}
