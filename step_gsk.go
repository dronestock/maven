package main

import (
	"context"
	"sync"

	"github.com/goexl/gox/field"
)

type stepGsk struct {
	*plugin
}

func newGskStep(plugin *plugin) *stepGsk {
	return &stepGsk{
		plugin: plugin,
	}
}

func (g *stepGsk) Runnable() bool {
	return !g.private()
}

func (g *stepGsk) Run(ctx context.Context) (err error) {
	wg := new(sync.WaitGroup)
	wg.Add(len(g.Repositories))
	for _, _repository := range g.Repositories {
		go g.gsk(ctx, _repository, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (g *stepGsk) gsk(_ context.Context, repository *repository, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	args := []any{
		"--server",
		g.Gpg.Server,
		"--username",
		repository.Username,
	}
	if ee := g.Command(gskExe).Args(args...).Dir(g.Source).Exec(); nil != ee {
		*err = ee
		g.Warn("生成密钥出错", field.New("repository", repository))
	}
}
