package main

const (
	exe         = `maven`
	homeEnv     = `HOME`
	pomFilename = `pom.xml`

	configMirrors    = `mirrors`
	configProperties = `properties`

	sourcesPath = `plugin[artifactId=maven-source-plugin]`
	docsPath    = `plugin[artifactId=maven-javadoc-plugin]`

	toIntFormat = `%d`
	xmlSpaces   = 2

	mavenDir         = `.m2`
	settingsFilename = `settings.xml`
	repositoryId     = `repository`

	xmlDeclare        = `version="1.0" encoding="UTF-8"`
	xmlSettingsXmlns  = `http://maven.apache.org/SETTINGS/1.0.0`
	xmlSettingsXsi    = `http://www.w3.org/2001/XMLSchema-instance`
	xmlSettingsSchema = `http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/settings-1.0.0.xsd`
	xmlProjectXmlns   = `http://maven.apache.org/POM/4.0.0`
	xmlProjectXsi     = `http://www.w3.org/2001/XMLSchema-instance`
	xmlProjectSchema  = `http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd`

	xmlPluginSourceArtifact = `maven-source-plugin`
	xmlPluginSources        = `attach-sources`
	xmlPluginPackage        = `package`
	xmlPluginJarNoFork      = `jar-no-fork`

	xmlPluginDocArtifact = `maven-javadoc-plugin`
	xmlPluginDocs        = `attach-javadocs`
	xmlPluginJar         = `jar`
)
