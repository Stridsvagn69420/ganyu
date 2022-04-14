package custom

import (
	"encoding/json"
	"path/filepath"

	"github.com/Stridsvagn69420/ganyu/utils"
)

var CustomCommandPath = filepath.Join(utils.GetHomeDir(), ".config/ganyu/custom.json")

type Custom struct {
	Desc string   `json:"desc"`
	Name string   `json:"name"`
	Root bool     `json:"root"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func ReadCustom(path string) ([]Custom, error) {
	data, err := utils.ReadFileByteArray(path)
	if err != nil {
		return nil, err
	}
	var custom []Custom
	err = json.Unmarshal(data, &custom)
	if err != nil {
		return nil, err
	}
	return custom, err
}
