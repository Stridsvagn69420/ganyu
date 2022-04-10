package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Choco(root bool) {
	utils.RunShell(root, "choco", "upgrade", "all")
}
