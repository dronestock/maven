package main

import (
	"context"

	"github.com/goexl/gox/args"
)

type stepPackage struct {
	*plugin
}

func newPackageStep(plugin *plugin) *stepPackage {
	return &stepPackage{
		plugin: plugin,
	}
}

func (p *stepPackage) Runnable() bool {
	return true
}

func (p *stepPackage) Run(_ context.Context) (err error) {
	builder := args.New().Build()
	// 清理
	if p.Clean {
		builder.Subcommand("clean")
	}

	// 测试
	if p.Test {
		builder.Subcommand("test")
	}

	// 打包
	builder.Subcommand("package")

	// 测试参数
	if !p.Test {
		builder.Arg("define", "maven.test.skip=true")
	}

	// 执行命令
	for _, repo := range p.Repositories {
		err = p.mvn(builder, repo)
		if nil != err {
			break
		}
	}

	return
}
