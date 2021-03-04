package main

import (
	"os"

	pb "AppFactory/api/webApp/v1"
	"AppFactory/internal/service"
	"AppFactory/pkg/config"
	mylog "AppFactory/pkg/log"

	"github.com/go-kratos/kratos/v2"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/spf13/cobra"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version  string
	confName string
	cmd      = &cobra.Command{
		Use:     "",
		Short:   "",
		Example: "",
		Run: func(cmd *cobra.Command, args []string) {
			// 组件初始化
			// cfg := config.InitConfig(confName)
			cfgYaml := config.InitConfigYaml(confName)
			// 日志初始化
			mylog.InitLogger(cfgYaml)
			logger := mylog.GetLogInstance()

			app, err := initApp(cfgYaml, logger)
			if err != nil {
				panic(err)
			}

			// start and wait for stop signal
			if err := app.Run(); err != nil {
				panic(err)
			}

		},
	}
)

func init() {
	cmd.PersistentFlags().StringVarP(&confName, "config", "c", "", "the name of TOML configuration file in config directory")

}

func newApp(logger *mylog.ZapLog, hs *http.Server, gs *grpc.Server, service *service.AppExcelService) *kratos.App {
	pb.RegisterAppExcelServer(gs, service)
	pb.RegisterAppExcelHTTPServer(hs, service)
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
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
