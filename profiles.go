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
	profiles := settings.SelectElement(keyProfiles)
	if nil == profiles {
		profiles = settings.CreateElement(keyProfiles)
	}
	profile := profiles.SelectElement(keyProfile)
	if nil == profile {
		profile = profiles.CreateElement(keyProfile)
	}
	profile.CreateElement(keyId).SetText(p.repositoryId(p.Repository.Release))
	profile.CreateElement(keyActivation).CreateElement(keyActivationDefault).SetText(xmlTrue)

	sign := profile.CreateElement(keyProperties)
	sign.CreateElement(keyGpgExecutable).SetText(xmlGpgExecutable)
	sign.CreateElement(keyGpgPassphrase).SetText(p.Password)
}
