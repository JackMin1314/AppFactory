// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// QueryExamMain is the golang structure for table query_exam_main.
type QueryExamMain struct {
    Id           uint64      `orm:"id,primary"    json:"id"`            //           
    ExamNum      string      `orm:"exam_num"      json:"exam_num"`      // 学生考号  
    StudentName  string      `orm:"student_name"  json:"student_name"`  // 学生姓名  
    ClassName    string      `orm:"class_name"    json:"class_name"`    // 学生班级  
    ChineseScore string      `orm:"chinese_score" json:"chinese_score"` // 语文成绩  
    MathScore    string      `orm:"math_score"    json:"math_score"`    // 数学成绩  
    EnglishScore string      `orm:"english_score" json:"english_score"` // 英语成绩  
    TotalScore   string      `orm:"total_score"   json:"total_score"`   // 学生总分  
    ClassRate    string      `orm:"class_rate"    json:"class_rate"`    // 班级排名  
    SchoolRate   string      `orm:"school_rate"   json:"school_rate"`   // 年级排名  
    StepRank     string      `orm:"step_rank"     json:"step_rank"`     // 进退名次  
    UploadDate   string      `orm:"upload_date"   json:"upload_date"`   // 上传日期  
    IsDeleted    string      `orm:"is_deleted"    json:"is_deleted"`    // 是否删除  
    DeleteTime   string      `orm:"delete_time"   json:"delete_time"`   // 删除时间  
    CreatedAt    *gtime.Time `orm:"created_at"    json:"created_at"`    //           
    UpdatedAt    *gtime.Time `orm:"updated_at"    json:"updated_at"`    //           
}