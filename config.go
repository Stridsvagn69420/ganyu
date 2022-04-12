package main

import (
	"encoding/json"
	"path/filepath"

	"github.com/Stridsvagn69420/ganyu/utils"
)

type Config struct {
	Sysupdate Sysupdate `json:"sysupdate"`
	Ytdl      []Ytdl    `json:"yt-dl"`
}

type Ytdl struct {
	Audio      string `json:"audio"`
	Video      string `json:"video"`
	AudioVideo string `json:"audio+video"`
	Website    string `json:"website"`
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
