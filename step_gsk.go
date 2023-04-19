package main

import (
	"context"
	"sync"

	"github.com/goexl/gox/args"
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
	return g.deploy() && !g.private()
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

	builder := args.New().Build()
	builder.Arg("server", g.Gpg.Server)
	builder.Arg("username", repository.Username)
	if _, ee := g.Command(g.Binary.Gsk).Args(builder.Build()).Dir(g.Source).Build().Exec(); nil != ee {
		*err = ee
		g.Warn("生成密钥出错", field.New("repository", repository))
	}
}
