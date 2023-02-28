package main

const (
	java  = "JAVA"
	certs = "/lib/security/cacerts"

	exe         = "mvn"
	gpgExe      = "gpg"
	gskExe      = "gsk"
	homeEnv     = "HOME"
	gpgHome     = ".gnupg"
	pomFilename = "pom.xml"

	randLength = 8
	xmlSpaces  = 2
	xmlTrue    = "true"
	xmlAll     = "*"
	xmlAlways  = "always"

	mavenHome           = ".m2"
	settingsFilename    = "settings.xml"
	mavenRepositoryHost = "https://s01.oss.sonatype.org"

	xmlDeclare        = `version="1.0" encoding="UTF-8"`
	xmlSettingsXmlns  = "http://maven.apache.org/SETTINGS/1.0.0"
	xmlSettingsXsi    = "http://www.w3.org/2001/XMLSchema-instance"
	xmlSettingsSchema = "http://maven.apache.org/SETTINGS/1.0.0 http://maven.apache.org/xsd/global-1.0.0.xsd"
	xmlProjectXmlns   = "http://maven.apache.org/POM/4.0.0"
	xmlProjectXsi     = "http://www.w3.org/2001/XMLSchema-instance"
	xmlProjectSchema  = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd"

	xmlMavenGroup = "org.apache.maven.plugins"
)
