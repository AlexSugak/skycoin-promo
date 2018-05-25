package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"
)

func testCommand() gcli.Command {
	name := "testCommand"
	// TODO: Remove it when new commands are implemented
	// Following statement is used to suppress linter warning
	_ = onCommandUsageError("suppress linter error")
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
				return gcli.ShowSubcommandHelp(c)
			}

			fmt.Println(testArg)

			return nil
		},
	}
}
