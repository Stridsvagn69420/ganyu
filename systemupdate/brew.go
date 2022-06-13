package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Brew(root bool) error {
	utils.RunShell(root, "brew", "update")
	return utils.RunShell(root, "brew", "upgrade")
}
