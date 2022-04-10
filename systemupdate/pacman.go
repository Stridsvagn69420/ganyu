package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Pacman() {
	utils.RunShell("pacman", "--noconfirm", "-Syu")
}

func Yay() {
	utils.RunShell("yay", "--noconfirm", "--answerclean", "None", "--answerdiff", "None", "-Syu")
}
