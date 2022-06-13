package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Apt(root bool) error {
	utils.RunShell(root, "apt", "update", "-y")
	return utils.RunShell(root, "apt", "upgrade", "-y")
}
