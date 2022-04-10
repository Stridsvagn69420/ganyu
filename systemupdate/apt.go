package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Apt(root bool) {
	utils.RunShell(root, "apt", "update", "-y")
	utils.RunShell(root, "apt", "upgrade", "-y")
}
