package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

// Still WIP, since Kagero is still under developement.
func Kagero(root bool) {
	utils.RunShell(root, "kagero", "update")
}

func Kaze(root bool) {
	utils.RunShell(root, "kaze", "update")
}
