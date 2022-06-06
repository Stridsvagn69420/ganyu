package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func pacman(root *bool) {
	utils.RunShell(*root, "pacman", "--noconfirm", "-Syu")
}

func yay() {
	utils.RunShell(false, "yay", "--noconfirm", "--answerclean", "None", "--answerdiff", "None", "-Syu")
}

func Arch(root bool) {
	if utils.CommandExists("yay") {
		yay()
	} else {
		pacman(&root)
	}
}
