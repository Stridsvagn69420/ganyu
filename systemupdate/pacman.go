package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func pacman(root *bool) error {
	return utils.RunShell(*root, "pacman", "--noconfirm", "-Syu")
}

func yay() error {
	return utils.RunShell(false, "yay", "--noconfirm", "--answerclean", "None", "--answerdiff", "None", "-Syu")
}

func Arch(root bool) error {
	if utils.CommandExists("yay") {
		return yay()
	} else {
		return pacman(&root)
	}
}
