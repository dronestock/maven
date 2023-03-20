package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/beevik/etree"
	"github.com/goexl/gfx"
	"github.com/goexl/gox/rand"
)

const (
	keySettings = "settings"

	keySettingsLocalRepository = "localRepository"
	localRepository            = "MAVEN_LOCAL_REPOSITORY"

	keySettingsGroups = "pluginGroups"
	keySettingsGroup  = "pluginGroup"
	xmlPlugins        = "org.sonatype.plugins"

	keyMirrors       = "mirrors"
	keyMirror        = "mirror"
	keyMirrorOf      = "mirrorOf"
	mirrorPathFormat = "mirror[url='%s']"

	serverPathFormat = "server[id='%s']"
	keyServers       = "servers"
	keyServer        = "server"
	keyUsername      = "username"
	keyPassword      = "password"

	keyProfiles          = "profiles"
	keyProfile           = "profile"
	keyActivation        = "activation"
	keyActivationDefault = "activeByDefault"
	keyGpgExecutable     = "gpg.executable"
	keyGpgPassphrase     = "gpg.passphrase"
	xmlGpgExecutable     = "gpg2"
	xmlGpgId             = "gpg"
)

type stepGlobal struct {
	*plugin
}

func newGlobalStep(plugin *plugin) *stepGlobal {
	return &stepGlobal{
		plugin: plugin,
	}
}

func (g *stepGlobal) Runnable() bool {
	return 0 != len(g.Repositories)
}

func (g *stepGlobal) Run(_ context.Context) (err error) {
	filename := filepath.Join(os.Getenv(homeEnv), mavenHome, settingsFilename)
	if _, exists := gfx.Exists(filename); !exists {
		if err = gfx.Create(filepath.Dir(filename), gfx.Dir()); nil != err {
			return
		}
		if err = gfx.Create(filename, gfx.File()); nil != err {
			return
		}
	}

	doc := etree.NewDocument()
	if err = doc.ReadFromFile(filename); nil != err {
		return
	}

	// 配置全局
	settings := g.settings(doc)
	// 本地仓库
	g.writeLocalRepository(settings)
	// 组信息
	g.writeGroups(settings)
	// 镜像
	g.writeMirrors(settings)
	// 仓库
	g.writeServers(settings)
	// 配置
	g.writeProfiles(settings)

	// 写入文件
	doc.Indent(xmlSpaces)
	err = doc.WriteToFile(filename)

	return
}

func (g *stepGlobal) writeLocalRepository(settings *etree.Element) {
	_repository := settings.SelectElement(keySettingsLocalRepository)
	if nil == _repository {
		_repository = settings.CreateElement(keySettingsLocalRepository)
		_repository.SetText(os.Getenv(localRepository))
	}
}

func (g *stepGlobal) writeGroups(settings *etree.Element) {
	groups := settings.CreateElement(keySettingsGroups)
	if nil != groups {
		group := groups.CreateElement(keySettingsGroup)
		group.CreateText(xmlPlugins)
	}
}

func (g *stepGlobal) writeMirrors(settings *etree.Element) {
	if 0 == len(g.mirrors()) {
		return
	}

	mirrors := settings.CreateElement(keyMirrors)
	for _, _mirror := range g.mirrors() {
		mirror := mirrors.FindElementPath(etree.MustCompilePath(fmt.Sprintf(mirrorPathFormat, _mirror)))
		if nil != mirror {
			mirrors.RemoveChildAt(mirror.Index())
		}

		id := rand.New().String().Length(randLength).Build().Generate()
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

func (g *stepGlobal) writeServers(settings *etree.Element) {
	servers := settings.CreateElement(keyServers)

	for _, _repository := range g.Repositories {
		// 写入正式服务器
		g.writeReleaseServer(servers, _repository)
		// 写入快照服务器
		g.writeSnapshotServer(servers, _repository)
	}
}

func (g *stepGlobal) writeReleaseServer(servers *etree.Element, repository *repository) {
	path := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, repository.releaseId()))
	release := servers.FindElementPath(path)
	if nil != release {
		servers.RemoveChildAt(release.Index())
	}
	release = servers.CreateElement(keyServer)
	release.CreateElement(keyId).SetText(repository.releaseId())
	release.CreateElement(keyUsername).SetText(repository.Username)
	release.CreateElement(keyPassword).SetText(repository.Password)
}

func (g *stepGlobal) writeSnapshotServer(servers *etree.Element, repository *repository) {
	path := etree.MustCompilePath(fmt.Sprintf(serverPathFormat, repository.snapshotId()))
	snapshot := servers.FindElementPath(path)
	if nil != snapshot {
		servers.RemoveChildAt(snapshot.Index())
	}
	snapshot = servers.CreateElement(keyServer)
	snapshot.CreateElement(keyId).SetText(repository.snapshotId())
	snapshot.CreateElement(keyUsername).SetText(repository.Username)
	snapshot.CreateElement(keyPassword).SetText(repository.Password)
}

func (g *stepGlobal) writeProfiles(settings *etree.Element) {
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
	properties.CreateElement(keyGpgPassphrase).SetText(g.passphrase())
}

func (g *stepGlobal) settings(doc *etree.Document) (settings *etree.Element) {
	settings = doc.SelectElement(keySettings)
	if nil == settings {
		doc.CreateProcInst(keyXml, xmlDeclare)
		settings = doc.CreateElement(keySettings)
		settings.CreateAttr(keyXmlns, xmlSettingsXmlns)
		settings.CreateAttr(keyXsi, xmlSettingsXsi)
		settings.CreateAttr(keySchema, xmlSettingsSchema)
	}

	return
}
