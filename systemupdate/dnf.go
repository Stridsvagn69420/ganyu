package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Dnf() {
	fedora("dnf")
}

func Yum() {
	fedora("yum")
}

func fedora(pkg string) {
	utils.RunShell(pkg, "check-update", "-y")
	utils.RunShell(pkg, "upgrade", "-y")
}
