package main

import (
	"context"

	"github.com/goexl/gox/args"
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
	builder := args.New().Build()
	builder.Subcommand("deploy")
	// 打印更多日志
	if d.Verbose {
		builder.Flag("X")
	}

	// 执行命令
	err = d.mvn(builder)

	return
}
