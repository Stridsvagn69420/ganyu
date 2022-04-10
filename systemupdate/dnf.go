package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func fedora(root *bool, pkg string) {
	utils.RunShell(*root, "check-update", "-y")
	utils.RunShell(*root, "upgrade", "-y")
}

func Yum(root *bool) {
	fedora(root, "yum")
}

func Dnf(root *bool) {
	fedora(root, "dnf")
}

func Fedora(root bool) {
	if utils.CommandExists("dnf") {
		Dnf(&root)
	} else {
		Yum(&root)
	}
}
