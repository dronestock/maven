package main

import (
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/goexl/gfx"
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

func (k *stepKeypair) Runnable() bool {
	return 0 != len(k.Repositories)
}

func (k *stepKeypair) Run(ctx context.Context) (err error) {
	// 删除原来的密钥目录
	if err = gfx.Delete(filepath.Join(os.Getenv(homeEnv), gpgHome)); nil != err {
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(k.Repositories))
	for _, _repository := range k.Repositories {
		go k.make(ctx, _repository, wg, &err)
	}

	// 等待所有任务执行完成
	wg.Wait()

	return
}

func (k *stepKeypair) make(_ context.Context, repository *repository, wg *sync.WaitGroup, err *error) {
	// 任何情况下，都必须调用完成方法
	defer wg.Done()

	args := []any{
		"--batch",
		"--passphrase",
		k.passphrase(),
		"--quick-gen-key",
		repository.Username,
		"default",
		"default",
		k.Gpg.Expire,
	}
	if ee := k.Command(gpgExe).Args(args...).Dir(k.Source).Exec(); nil != ee {
		*err = ee
		k.Warn("生成密钥出错", field.New("repository", repository))
	}
}
