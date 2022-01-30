package main

import (
	`fmt`
	`net/url`

	`github.com/beevik/etree`
	`github.com/storezhang/gox`
)

const (
	keyMirrors  = `mirrors`
	keyMirror   = `mirror`
	keyMirrorOf = `mirrorOf`

	mirrorPathFormat = `mirror[url='%s']`
)

func (p *plugin) mirrors(settings *etree.Element) () {
	if 0 == len(p.Mirrors) {
		return
	}

	mirrors := settings.CreateElement(keyMirrors)
	for _, _mirror := range p._mirrors() {
		mirror := mirrors.FindElementPath(etree.MustCompilePath(fmt.Sprintf(mirrorPathFormat, _mirror)))
		if nil != mirror {
			mirrors.RemoveChildAt(mirror.Index())
		}

		id := gox.RandString(randLength)
		if host, err := url.Parse(_mirror); nil == err {
			id = host.Hostname()
		}
		mirror = mirrors.CreateElement(keyMirror)
		mirror.CreateElement(keyId).CreateText(id)
		mirror.CreateElement(keyMirrorOf).CreateText(xmlAll)
		mirror.CreateElement(keyName).CreateText(id)
		mirror.CreateElement(keyUrl).CreateText(_mirror)
	}
}
