package main

import (
	`os`

	`github.com/beevik/etree`
)

const (
	keySettingsLocalRepository = `localRepository`
	localRepository            = `MAVEN_LOCAL_REPOSITORY`
)

func (p *plugin) localRepository(settings *etree.Element) {
	_repository := settings.SelectElement(keySettingsLocalRepository)
	if nil == _repository {
		_repository = settings.CreateElement(keySettingsLocalRepository)
		_repository.SetText(os.Getenv(localRepository))
	}
}
