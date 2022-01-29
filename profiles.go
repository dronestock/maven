package main

import (
	`github.com/beevik/etree`
)

const (
	keyProfiles          = `profiles`
	keyProfile           = `profile`
	keyActivation        = `activation`
	keyActivationDefault = `activeByDefault`
	keyGpgExecutable     = `gpg.executable`
	keyGpgPassphrase     = `gpg.passphrase`

	xmlGpgExecutable = `gpg2`
)

func (p *plugin) profiles(settings *etree.Element) () {
	profiles := settings.CreateElement(keyProfiles)
	profile := profiles.CreateElement(keyProfile)
	profile.CreateElement(keyId).SetText(p.repositoryId())
	profile.CreateElement(keyActivation).CreateElement(keyActivationDefault).SetText(xmlTrue)

	properties := profile.CreateElement(keyProperties)
	properties.CreateElement(keyGpgExecutable).SetText(xmlGpgExecutable)
	if `` != p.GpgPassphrase {
		properties.CreateElement(keyGpgPassphrase).SetText(p.GpgPassphrase)
	}
}
