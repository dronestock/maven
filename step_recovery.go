package main

import (
	"context"
	"os"
)

type stepRecovery struct {
	*plugin
}

func newRecoveryStep(plugin *plugin) *stepRecovery {
	return &stepRecovery{
		plugin: plugin,
	}
}

func (r *stepRecovery) Runnable() bool {
	return true
}

func (r *stepRecovery) Run(_ context.Context) error {
	return os.WriteFile(r.original, r.content, r.mode)
}
