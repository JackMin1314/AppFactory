package data

import (
	pb "AppFactory/api/webApp/v1"
	"AppFactory/internal/biz"
	"AppFactory/internal/data/ent"
	"AppFactory/internal/data/ent/student"
	"AppFactory/pkg/log"
	"context"
	"encoding/json"
	"errors"
)

type AppExcelImplRepo struct {
	data *Data
	log  *log.ZapLog
}

func NewAppExcelImplRepo(data *Data, logger *log.ZapLog) biz.AppExcelRepo {
	return &AppExcelImplRepo{
		data: data,
		log:  logger,
	}
}
func (ap *AppExcelImplRepo) QueryScoreMain(ctx context.Context, stu *biz.AppExcel) (*pb.GetStudentReply, error) {
	// 先查redis
	val, _ := ap.data.RDGetStruct(ctx, FormatRedisKey(stu))
	if len(val) > 0 {
		reply := new(pb.GetStudentReply)
		if err := json.Unmarshal(val, reply); err == nil {
			return reply, nil
		}
	}
	// 查库添加redis
	entStu, err := ap.data.db.Student.Query().Where(student.ExamNum(stu.ExamNum), student.StudentName(stu.StudentName)).Only(ctx)
	if err != nil {
		if _, ok := err.(*ent.NotFoundError); ok {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	stuReply := &pb.GetStudentReply{
		ExamNum:      entStu.ExamNum,
		StudentName:  entStu.StudentName,
		ClassName:    entStu.ClassName,
		ChineseScore: entStu.ChineseScore,
		MathScore:    entStu.MathScore,
		EnglishScore: entStu.EnglishScore,
		TotalScore:   entStu.TotalScore,
		ClassRate:    entStu.ClassRate,
		SchoolRate:   entStu.SchoolRate,
		StepRank:     entStu.StepRank,
		UploadDate:   entStu.UploadDate,
		IsDeleted:    entStu.IsDeleted,
		DeleteTime:   entStu.DeleteTime,
	}
	ap.data.RDSetStruct(ctx, FormatRedisKey(stu), stuReply)

	return stuReply, nil
}
func (ap *AppExcelImplRepo) QueryScoreMajor(ctx context.Context, stu *biz.AppExcel) (*pb.GetStudentReply, error) {
	return nil, nil
}
