package main

import (
	"encoding/json"
	"path/filepath"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/ganyu/ytdl"
)

type Config struct {
	RPC       bool        `json:"rpc"`
	Sysupdate Sysupdate   `json:"sysupdate"`
	Ytdl      []ytdl.Ytdl `json:"yt-dl"`
	Git       GitRepos    `json:"git"`
}

type GitRepos struct {
	BaseDir string   `json:"base"`
	Repos   []string `json:"repos"`
}

type Sysupdate struct {
	Root     bool `json:"root"`
	CrossPkg bool `json:"crosspkg"`
}

func config() (Config, error) {
	var config Config
	configpath := filepath.Join(utils.GetHomeDir(), ".config/ganyu/config.json")
	data, err := utils.ReadFileByteArray(configpath)
	if err != nil {
		return Config{}, err
	}
	json.Unmarshal(data, &config)
	return config, err
}
