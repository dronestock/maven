package main

var (
	defaultProperties = map[string]string{
		"maven.compiler.source": "1.8",
		"maven.compiler.target": "1.8",
		"java.version":          "1.8",
	}
	defaultDefines = map[string]string{
		"maven.wagon.http.ssl.insecure":              "true",
		"maven.wagon.http.ssl.allowall":              "true",
		"maven.wagon.http.ssl.ignore.validity.dates": "true",
	}
	defaultMirrors []string
)
