package main

import (
	`os`

	`github.com/beevik/etree`
)

const (
	keySettingsLocalRepository = `localRepository`
	localRepository            = `MAVEN_LOCAL_REPOSITORY`
)

func (p *plugin) repository(settings *etree.Element) {
	repository := settings.SelectElement(keySettingsLocalRepository)
	if nil == repository {
		repository = settings.CreateElement(keySettingsLocalRepository)
		repository.SetText(os.Getenv(localRepository))
	}
}
