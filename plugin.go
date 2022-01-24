package main

import (
	`github.com/dronestock/drone`
)

type plugin struct {
	config *config
	envs   []string
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: new(config),
		envs:   make([]string, 0),
	}
}

func (p *plugin) Configuration() drone.Configuration {
	return p.config
}

func (p *plugin) Steps() []*drone.Step {
	return []*drone.Step{
		drone.NewStep(p.test, drone.Name(`测试`)),
	}
}
