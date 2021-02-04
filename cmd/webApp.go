package main

import (
	webapp "AppFactory/app/webApp"
	"os"

	"github.com/spf13/cobra"
)

var (
	confName string
	cmd      = &cobra.Command{
		Use:     "",
		Short:   "",
		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			// 放在最前面，部分组件初始化
			webapp.Setup(confName)
			// 服务启动
			webapp.RunExec()
		},
	}
)

func init() {
	cmd.PersistentFlags().StringVarP(&confName, "config", "c", "", "the name of TOML configuration file in config directory")

}

func main() {
	Execute()
}

// Execute cobra execute function
func Execute() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
