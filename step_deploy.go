package main

import (
	"context"
)

type stepDeploy struct {
	*plugin
}

func newDeployStep(plugin *plugin) *stepDeploy {
	return &stepDeploy{
		plugin: plugin,
	}
}

func (d *stepDeploy) Runnable() bool {
	return 0 != len(d.Repositories)
}

func (d *stepDeploy) Run(ctx context.Context) (err error) {
	args := []any{
		"deploy",
	}
	// 打印更多日志
	if d.Verbose {
		args = append(args, "-X")
	}

	// 执行命令
	err = d.mvn(args...)

	return
}
