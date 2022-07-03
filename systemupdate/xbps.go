package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Void(root bool) error {
	utils.RunShell(root, "xbps-install", "-u", "xbps")
	return utils.RunShell(root, "xbps-install", "-Su")
}
