package main

import (
	`github.com/dronestock/drone`
)

func (p *plugin) maven(folder string, args ...string) (err error) {
	return p.Exec(exe, drone.Args(args...), drone.Dir(folder))
}
