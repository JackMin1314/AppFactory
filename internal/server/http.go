package server

import (
	pb "AppFactory/api/webApp/v1"
	"AppFactory/internal/service"
	"AppFactory/pkg/config"
	"time"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *config.ConfigYaml, appsrv *service.AppExcelService) *http.Server {
	var opts = []http.ServerOption{}

	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != 0 {
		opts = append(opts, http.Timeout(time.Duration(c.Server.Http.Timeout)*time.Second))
	}
	srv := http.NewServer(opts...)
	m := http.Middleware(
		middleware.Chain(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(),
		),
	)
	srv.HandlePrefix("/", pb.NewAppExcelHandler(appsrv, m))

	return srv
}
