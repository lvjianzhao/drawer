package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"server/core"
	"server/global"
	"server/initialize"
)

var (
	config   string
	StartCmd = &cobra.Command{
		Use:     "server",
		Short:   "Start API server",
		Example: "ferry server config/settings.yml",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "Start server with provided configuration file")
}

func usage() {
	usageStr := `starting api server`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	// 1. 读取配置
	global.TD27_VP = core.Viper(config) // 初始化viper
}

func run() error {
	global.TD27_LOG = core.Zap() // 初始化zap日志
	zap.ReplaceGlobals(global.TD27_LOG)
	global.TD27_DB = initialize.Gorm() // gorm连接数据库

	if global.TD27_DB == nil {
		global.TD27_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.TD27_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.TD27_DB.DB()
		defer db.Close()
	}
	core.RunServer()
	return nil
}
