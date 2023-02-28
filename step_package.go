package main

import (
	"context"
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
	args := make([]any, 0)

	// 清理
	if p.Clean {
		args = append(args, "clean")
	}

	// 测试
	if p.Test {
		args = append(args, "test")
	}

	// 打包
	args = append(args, "package")

	// 测试参数
	if !p.Test {
		args = append(args, "--define", "maven.skip.test=true")
	}

	// 打印更多日志
	if p.Verbose {
		args = append(args, "-X")
	}

	// 执行命令
	err = p.mvn(args...)

	return
}
