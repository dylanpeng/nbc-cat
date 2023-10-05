package cmd

import (
	"cat/common"
	"cat/internal/admin/config"
	"cat/internal/admin/router"
	"encoding/json"
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
	_ = startCmd.MarkFlagRequired("config")
}

func (app *Application) Init(env svc.Environment) error {
	if err := config.Init(*cfgFile); err != nil {
		fmt.Printf("config init err: %s", err)
		panic(err)
	}

	conf := config.GetConfig()
	b, _ := json.Marshal(conf)
	fmt.Println("config: ", string(b))

	// 初始化log
	if err := common.InitLogger(conf.Log); err != nil {
		fmt.Println("logger init err: ", err)
		panic(err)
	}

	// 初始化redis
	common.InitCache()

	// 初始化mysql
	if err := common.InitDB(conf.DB); err != nil {
		fmt.Printf("mysql init err: %s", err)
		panic(err)
	}

	fmt.Println("init")
	common.Logger.Infof("admin init")

	return nil
}

func (app *Application) Start() error {
	// http服务
	common.InitHttpServer(router.Router)
	httpConf := config.GetConfig().Http

	fmt.Println("start")
	common.Logger.Infof("admin start at %s:%d", httpConf.Host, httpConf.Port)
	return nil
}

func (app *Application) Stop() error {
	// do stop

	fmt.Println("stop")
	common.Logger.Infof("admin stop")
	_ = common.Logger.Sync()
	_ = common.Logger.Close()
	return nil
}
