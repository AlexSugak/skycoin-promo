package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"
)

// App represents CLI app
type App struct {
	gcli.App
}

// NewApp creates a new instance of the App
func NewApp() *App {
	gcliApp := gcli.NewApp()

	commands := []gcli.Command{
		testCommand(),
	}

	gcliApp.Commands = commands

	return &App{
		App: *gcliApp,
	}
}

// Run starts the app
func (app *App) Run(args []string) error {
	return app.App.Run(args)
}

func onCommandUsageError(command string) gcli.OnUsageErrorFunc {
	return func(c *gcli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "Error: %v\n\n", err)
		return gcli.ShowCommandHelp(c, command)
	}
}
