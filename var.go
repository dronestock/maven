package main

var (
	defaultProperties = map[string]string{
		`maven.compiler.source`: `17`,
		`maven.compiler.target`: `17`,
		`java.version`:          `17`,
	}
	defaultDefines = map[string]string{
		`maven.wagon.http.ssl.insecure`:              `true`,
		`maven.wagon.http.ssl.allowall`:              `true`,
		`maven.wagon.http.ssl.ignore.validity.dates`: `true`,
	}
	defaultMirrors = []string{
		`https://mirrors.cloud.tencent.com/nexus/repository/maven-public/`,
		`https://maven.aliyun.com/repository/public/`,
		`https://mirrors.163.com/maven/repository/maven-public/`,
		`https://mirrors.huaweicloud.com/repository/maven/`,
	}
)
