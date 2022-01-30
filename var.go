package main

var (
	defaultProperties = map[string]string{
		`maven.compiler.source`: `17`,
		`maven.compiler.target`: `17`,
		`java.version`:          `17`,
	}
	defaultMirrors = []string{
		`https://mirrors.huaweicloud.com/repository/maven/`,
		`https://maven.aliyun.com/repository/public/`,
		`https://mirrors.163.com/maven/repository/maven-public/`,
		`https://mirrors.cloud.tencent.com/nexus/repository/maven-public/`,
	}
)
