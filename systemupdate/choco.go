package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Choco(root bool) error {
	return utils.RunShell(root, "choco", "upgrade", "all")
}
