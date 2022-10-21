package main

func (p *plugin) deploy() (undo bool, err error) {
	if undo = 0 == len(p.Repositories); undo {
		return
	}

	args := []any{
		`deploy`,
	}
	// 打印更多日志
	if p.Verbose {
		args = append(args, `-X`)
	}

	// 执行命令
	err = p.mvn(args...)

	return
}
