package main

type repository struct {
	// 正式仓库
	Release string `default:"https://s01.oss.sonatype.org/service/local/staging/deploy/maven2"`
	// 快照仓库
	Snapshot string `default:"https://s01.oss.sonatype.org/content/repositories/snapshots"`
}
