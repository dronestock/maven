package main

import (
	"strings"

	"github.com/beevik/etree"
	"github.com/goexl/gox"
)

const (
	keyProfiles          = `profiles`
	keyProfile           = `profile`
	keyActivation        = `activation`
	keyActivationDefault = `activeByDefault`
	keyGpgExecutable     = `gpg.executable`
	keyGpgPassphrase     = `gpg.passphrase`

	xmlGpgExecutable = `gpg2`
	xmlGpgId         = `gpg`
)

func (p *plugin) writeProfiles(settings *etree.Element) {
	profiles := settings.SelectElement(keyProfiles)
	if nil == profiles {
		profiles = settings.CreateElement(keyProfiles)
	}

	profile := profiles.SelectElement(keyProfile)
	if nil == profile {
		profile = profiles.CreateElement(keyProfile)
	}
	profile.CreateElement(keyId).SetText(xmlGpgId)
	profile.CreateElement(keyActivation).CreateElement(keyActivationDefault).SetText(xmlTrue)

	properties := profile.CreateElement(keyProperties)
	properties.CreateElement(keyGpgExecutable).SetText(xmlGpgExecutable)
	passphrase := p.Gpg.Passphrase
	if `` == strings.TrimSpace(passphrase) {
		passphrase = gox.RandString(randLength)
	}
	properties.CreateElement(keyGpgPassphrase).SetText(passphrase)
}
