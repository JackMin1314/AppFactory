package biz

import (
	pb "AppFactory/api/webApp/v1"
	"AppFactory/pkg/log"
	"context"
	"errors"
)

type AppExcel struct {
	ExamNum     string `json:"exam_num,omitempty" validate:"required"`
	StudentName string `json:"student_name,omitempty" validate:"required"`
}

type AppExcelRepo interface {
	QueryScoreMain(context.Context, *AppExcel) (*pb.GetStudentReply, error)
	// UpdateScoreMain(context.Context, *AppExcel) (*pb.GetStudentReply, error)
	// DeleteScoreMain(context.Context, *AppExcel) (*pb.GetStudentReply, error)
	InsertScoreMain(context.Context, *AppExcel) (*pb.GetStudentReply, error)
	// InsertScoreAllMain(context.Context, []*AppExcel) (*pb.GetStudentReply, error)
}
type AppExcelUsecase struct {
	repo   AppExcelRepo
	logger *log.ZapLog
}

func NewAppExcelUsecase(bizRepo AppExcelRepo, logger *log.ZapLog) *AppExcelUsecase {
	return &AppExcelUsecase{
		repo:   bizRepo,
		logger: logger,
	}
}

func (uc *AppExcelUsecase) QueryMain(ctx context.Context, appexcel *AppExcel) (*pb.GetStudentReply, error) {
	if appexcel == nil || appexcel.ExamNum == "" || appexcel.StudentName == "" {
		return nil, errors.New("学生学号或姓名不能为空！")
	}
	uc.logger.Infof("收到查询主科目的学生学号[%s]和姓名[%s]", appexcel.ExamNum, appexcel.StudentName)
	stReply, err := uc.repo.QueryScoreMain(ctx, appexcel)
	if err != nil {
		return nil, err
	}
	return stReply, nil
}

func (uc *AppExcelUsecase) InsertMain(ctx context.Context, appexcel *AppExcel) (*pb.GetStudentReply, error) {
	if appexcel == nil || appexcel.ExamNum == "" || appexcel.StudentName == "" {
		return nil, errors.New("学生学号或姓名不能为空！")
	}
	uc.logger.Infof("收到查询核心科目的学生学号[%s]和姓名[%s]", appexcel.ExamNum, appexcel.StudentName)
	stReply, err := uc.repo.InsertScoreMain(ctx, appexcel)
	if err != nil {
		return nil, err
	}
	return stReply, nil
}
