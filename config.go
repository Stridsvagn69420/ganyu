package main

import (
	"encoding/json"
	"path/filepath"

	"github.com/Stridsvagn69420/ganyu/utils"
)

type Config struct {
	Sysupdate Sysupdate `json:"sysupdate"`
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
