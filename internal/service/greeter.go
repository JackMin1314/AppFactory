package service

import (
	"context"
	"errors"

	v1 "AppFactory/api/webApp/v1"
	"AppFactory/internal/biz"
	// mylog "AppFactory/pkg/log"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "jack" {
		s.log.Errorf("received name [%s] is forbidden", in.GetName())
		return &v1.HelloReply{Message: "you are jack,forbidden " + in.GetName()}, errors.New("user not allowed in ")
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}
