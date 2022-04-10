package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Snap(root bool) {
	utils.RunShell(root, "snap", "refresh")
}
