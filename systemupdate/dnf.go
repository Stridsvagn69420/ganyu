package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func fedora(root *bool, pkg string) error {
	utils.RunShell(*root, pkg, "check-update", "-y")
	return utils.RunShell(*root, pkg, "upgrade", "-y")
}

func Yum(root *bool) error {
	return fedora(root, "yum")
}

func Dnf(root *bool) error {
	return fedora(root, "dnf")
}

func Fedora(root bool) error {
	if utils.CommandExists("dnf") {
		return Dnf(&root)
	} else {
		return Yum(&root)
	}
}
