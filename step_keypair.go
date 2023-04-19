package main

import (
	"context"
	"sync"

	"github.com/goexl/gfx"
	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

type stepKeypair struct {
	*plugin
}

func newKeypairStep(plugin *plugin) *stepKeypair {
	return &stepKeypair{
		plugin: plugin,
	}
}

func (k *stepKeypair) Runnable() (runnable bool) {
	return k.deploy()
}

func (k *stepKeypair) Run(ctx context.Context) (err error) {
	// 删除原来的密钥目录
	if err = gfx.Delete(k.Home(gpgHome)); nil != err {
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(k.Repositories))
	for _, repo := range k.Repositories {
		go k.make(ctx, repo, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (k *stepKeypair) make(_ context.Context, repository *repository, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	builder := args.New().Build()
	builder.Flag("batch")
	builder.Option("passphrase", k.passphrase())
	builder.Option("quick-gen-key", repository.Username)
	builder.Add("default", "default", k.Gpg.Expire)
	if _, ee := k.Command(k.Binary.Gpg).Args(builder.Build()).Dir(k.Source).Build().Exec(); nil != ee {
		*err = ee
		k.Warn("生成密钥出错", field.New("repository", repository))
	}
}
