package main

import (
	`fmt`

	`github.com/beevik/etree`
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
	count := 1
	for _, url := range p._mirrors() {
		mirror := mirrors.FindElementPath(etree.MustCompilePath(fmt.Sprintf(mirrorPathFormat, url)))
		if nil != mirror {
			mirrors.RemoveChildAt(mirror.Index())
		}

		mirror = mirrors.CreateElement(keyMirror)
		mirror.CreateElement(keyId).CreateText(fmt.Sprintf(toIntFormat, count))
		mirror.CreateElement(keyMirrorOf).CreateText(xmlCentral)
		mirror.CreateElement(keyName).CreateText(fmt.Sprintf(toIntFormat, count))
		mirror.CreateElement(keyUrl).CreateText(url)

		count++
	}
}
