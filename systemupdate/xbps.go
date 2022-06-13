package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Void(root bool) error {
	return utils.RunShell(root, "xbps-install", "-Su")
}
