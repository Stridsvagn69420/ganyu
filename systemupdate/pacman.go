package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func pacman(root *bool) {
	utils.RunShell(*root, "pacman", "--noconfirm", "-Syu")
}

func yay(root *bool) {
	utils.RunShell(*root, "yay", "--noconfirm", "--answerclean", "None", "--answerdiff", "None", "-Syu")
}

func Arch(root bool) {
	if utils.CommandExists("yay") {
		yay(&root)
	} else {
		pacman(&root)
	}
}
