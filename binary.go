package main

type binary struct {
	Maven string `default:"mvn" json:"maven"`
	Gpg   string `default:"gpg" json:"gpg"`
	Gsk   string `default:"gsk" json:"gsk"`
}
