package custom

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
)

func RunCustom(cmd Custom) error {
	return utils.RunShell(cmd.Root, cmd.Cmd, cmd.Args...)
}

func FindCustom(name string, commands []Custom) (Custom, bool) {
	for _, cmd := range commands {
		if os.Args[1] == cmd.Name {
			return cmd, true
		}
	}
	return Custom{}, false
}
