package main

import (
	"os"

	pb "AppFactory/api/webApp/v1"
	"AppFactory/internal/conf"
	"AppFactory/internal/service"
	"AppFactory/pkg/config"
	mylog "AppFactory/pkg/log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
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
			cfg := config.InitConfig(confName)
			cfgYaml := config.InitConfigYaml(confName)
			// 日志初始化
			mylog.InitLogger(cfgYaml)
			logger := mylog.GetLogInstance()
			logger.Info("cobra cmd init ")

			var bc conf.Bootstrap
			cfg.Scan(&bc)

			app, err := initApp(bc.Server, bc.Data, logger)
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

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, greeter *service.GreeterService) *kratos.App {
	pb.RegisterGreeterServer(gs, greeter)
	pb.RegisterGreeterHTTPServer(hs, greeter)
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
