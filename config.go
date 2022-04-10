package main

type Config struct {
	Sysupdate Sysupdate `json:"sysupdate"`
}

type Sysupdate struct {
	Root     bool `json:"root"`
	CrossPkg bool `json:"crosspkg"`
}

func config() Config {
	return Config{}
}
