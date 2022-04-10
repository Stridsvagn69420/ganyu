package utils

import (
	"io/ioutil"
	"os"
	"runtime"
)

func ReadFileByteArray(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ReadFileString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func GetHomeDir() string {
	var homedir string
	homedir, err := os.UserHomeDir()
	if err != nil {
		if runtime.GOOS == "windows" {
			homedir = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
			if homedir == "" {
				homedir = os.Getenv("USERPROFILE")
			}
		} else {
			homedir = os.Getenv("HOME")
		}
	}
	return homedir
}
