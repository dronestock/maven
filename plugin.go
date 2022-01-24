package main

import (
	`github.com/dronestock/drone`
)

type plugin struct {
	config *config
}

func newPlugin() drone.Plugin {
	return &plugin{
		config: new(config),
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
