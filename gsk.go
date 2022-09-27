package main

import (
	"github.com/dronestock/drone"
)

func (p *plugin) gsk() (undo bool, err error) {
	if undo = 0 == len(p.Repositories); undo {
		return
	}

	for _, _repository := range p.Repositories {
		args := []interface{}{
			`--server`,
			p.Gpg.Server,
			`--username`,
			_repository.Username,
		}
		err = p.Exec(gskExe, drone.Args(args...), drone.Dir(p.Source))
		if nil != err {
			return
		}
	}

	return
}
