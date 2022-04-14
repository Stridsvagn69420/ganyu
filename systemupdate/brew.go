package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Brew(root bool) {
	utils.RunShell(root, "brew", "update")
	utils.RunShell(root, "brew", "upgrade")
}
