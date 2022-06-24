package utils

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func RunShell(root bool, cmd string, args ...string) error {
	var c *exec.Cmd
	if root {
		if runtime.GOOS == "windows" {
			if CommandExists("wsudox") {
				c = exec.Command("wsudox", append([]string{cmd}, args...)...)
			} else {
				c = exec.Command(cmd, args...)
			}
		} else {
			c = exec.Command("sudo", append([]string{cmd}, args...)...)
		}
	} else {
		c = exec.Command(cmd, args...)
	}
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}

func RunShellCwd(directory string, root bool, cmd string, args ...string) error {
	var c *exec.Cmd
	if root {
		if runtime.GOOS == "windows" {
			if CommandExists("wsudox") {
				c = exec.Command("wsudox", append([]string{cmd}, args...)...)
			} else {
				c = exec.Command(cmd, args...)
			}
		} else {
			c = exec.Command("sudo", append([]string{cmd}, args...)...)
		}
	} else {
		c = exec.Command(cmd, args...)
	}
	if pathdir, err := filepath.Abs(directory); err != nil {
		c.Dir = directory
	} else {
		c.Dir = pathdir
	}
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	return c.Run()
}

func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
