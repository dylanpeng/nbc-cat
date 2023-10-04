package cmd

import (
	"fmt"
	"github.com/judwhite/go-svc"
	"github.com/spf13/cobra"
	"syscall"
)

type Application struct {
}

var cfgFile *string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the api server",
	Long: `usage example:
	server(.exe) start -c apollo.json
	start the server`,
	Run: func(cmd *cobra.Command, args []string) {
		app := &Application{}
		if err := svc.Run(app, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	cfgFile = startCmd.Flags().StringP("config", "c", "", "config file (required)")
	startCmd.MarkFlagRequired("config")
}

func (app *Application) Init(env svc.Environment) error {
	// do init
	fmt.Println("init")
	return nil
}

func (app *Application) Start() error {
	// do start
	fmt.Println("start")
	return nil
}

func (app *Application) Stop() error {
	// do stop
	fmt.Println("stop")
	return nil
}
