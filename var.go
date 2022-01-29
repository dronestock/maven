package main

var (
	defaultProperties = map[string]string{
		`maven.compiler.source`: `17`,
		`maven.compiler.target`: `17`,
		`java.version`:          `17`,
	}
	defaultMirrors = []string{
		`https://mirrors.163.com/maven/repository/maven-public/`,
	}
)
