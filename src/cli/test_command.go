package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"
)

func testCommand() gcli.Command {
	name := "testCommand"
	return gcli.Command{
		Name:         name,
		Usage:        "Test command of the cli (works as echo)",
		ArgsUsage:    "[test_arg]",
		Description:  fmt.Sprintf(`Test command of the cli (works as echo)`),
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) error {
			// get test arg
			testArg := c.Args().First()
			if testArg == "" {
				gcli.ShowSubcommandHelp(c)
				return nil
			}

			fmt.Println(testArg)

			return nil
		},
	}
}
