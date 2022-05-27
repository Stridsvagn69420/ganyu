package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
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

func IsInGopath(location string) bool {
	gopath := os.Getenv("GOPATH")
	appname := "ganyu"
	if runtime.GOOS == "windows" {
		appname = "ganyu.exe"
	}
	target := filepath.Join(gopath, "bin", appname)
	return location == target
}

func ReadFiles(location string) []string {
	files, err := ioutil.ReadDir(location)
	if err != nil {
		return make([]string, 0)
	}

	var filelist []string
	for _, file := range files {
		if !file.IsDir() {
			filelist = append(filelist, file.Name())
		}
	}
	return filelist
}

func DirExists(dir string) bool {
	_, err := os.Stat(dir)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
