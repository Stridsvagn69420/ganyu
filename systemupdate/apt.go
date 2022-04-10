package systemupdate

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Apt() {
	utils.RunShell("apt", "update", "-y")
	utils.RunShell("apt", "upgrade", "-y")
}
