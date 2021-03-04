package service

import (
	"context"
	"errors"

	pb "AppFactory/api/webApp/v1"
	"AppFactory/internal/biz"
	"AppFactory/pkg/log"
)

type AppExcelService struct {
	pb.UnimplementedAppExcelServer
	logger *log.ZapLog
	uc     *biz.AppExcelUsecase
}

func NewAppExcelService(logger *log.ZapLog, uc *biz.AppExcelUsecase) *AppExcelService {
	return &AppExcelService{
		logger: logger,
		uc:     uc,
	}
}

func (s *AppExcelService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	s.logger.Infof("[%s]调用sayhello方法", req.GetName())
	if req.GetName() == "jack" {
		s.logger.Errorf("received name [%s] is forbidden", req.GetName())
		return &pb.HelloReply{Message: "you are not allowed"}, errors.New("user is forbidden")
	}
	return &pb.HelloReply{Message: "hello,you call sayhello " + req.GetName()}, nil
}
func (s *AppExcelService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu := new(biz.AppExcel)
	stu.ExamNum = req.GetExamNum()
	stu.StudentName = req.GetStudentName()
	s.logger.Infof("调用GetStudent %+v", stu)
	stuReply, err := s.uc.QueryMain(ctx, stu)
	if err != nil {
		s.logger.Errorf("调用GetStudent异常%s", err)
		return nil, err
	}
	return stuReply, nil
}
