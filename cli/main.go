package main

import (
	cmd "depHooster/cli/cmds"
	uti "depHooster/cli/utils"
	"os"
)

func main() {
	userArgs := os.Args
	for _, arg := range userArgs {
		switch arg {
		case "-cl":
			uti.Show("[CONFIGURACAO LOCAL]")
			cmd.LocalConfiguration()
		case "-b":
			uti.Show("[BUILD]")
			cmd.Build()
		case "-h":
			uti.Show("[AJUDA]")
			cmd.Help()
		}
	}
}
