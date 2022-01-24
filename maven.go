package main

import (
	`github.com/storezhang/gex`
	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

func (p *plugin) maven(folder string, logger simaqian.Logger, args ...string) (err error) {
	fields := gox.Fields{
		field.String(`exe`, exe),
		field.Strings(`args`, args...),
		field.Bool(`verbose`, p.config.Verbose),
		field.Bool(`debug`, p.config.Debug),
	}
	// 记录日志
	logger.Info(`开始执行Maven命令`, fields...)

	options := gex.NewOptions(gex.Args(args...), gex.Dir(folder))
	if !p.config.Debug {
		options = append(options, gex.Quiet())
	}
	if _, err = gex.Run(exe, options...); nil != err {
		logger.Error(`执行Maven命令出错`, fields.Connect(field.Error(err))...)
	} else {
		logger.Info(`执行Maven命令成功`, fields...)
	}

	return
}
