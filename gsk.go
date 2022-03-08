package main

import (
	`strings`

	`github.com/dronestock/drone`
)

func (p *plugin) gsk() (undo bool, err error) {
	if undo = `` == strings.TrimSpace(p.Username) || `` == strings.TrimSpace(p.Password); undo {
		return
	}

	args := []interface{}{
		`--server`,
		p.Gpg.Server,
		`--username`,
		p.Username,
	}
	err = p.Exec(gskExe, drone.Args(args...), drone.Dir(p.Folder))

	return
}
