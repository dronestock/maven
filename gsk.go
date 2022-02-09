package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) gsk() (undo bool, err error) {
	args := []interface{}{
		`--server`,
		p.GpgServer,
		`--username`,
		p.Username,
	}
	err = p.Exec(gskExe, drone.Args(args...), drone.Dir(p.Folder))

	return
}
