// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"AppFactory/internal/dao/internal"
	"AppFactory/internal/model"
	"time"

	"github.com/gogf/gf/frame/g"
	"go.uber.org/zap"
)

// queryExamMainDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type queryExamMainDao struct {
	*internal.QueryExamMainDao
}

var (
	// QueryExamMain is globally public accessible object for table query_exam_main operations.
	QueryExamMain = &queryExamMainDao{
		internal.QueryExamMain,
	}
)

// Fill with you ideas below.

// InsertTableAll 带事务写入query_exam_main数据库逻辑; datasAll需要去除表头字段
func (QEM *queryExamMainDao) InsertTableAll(logger *zap.SugaredLogger, datasAll [][]string) error {
	db := g.DB()
	tx, err := db.Begin()
	if err != nil {
		tx.Rollback()
		logger.Errorf("开启事务失败[%s]", err)
		return err
	}
	QueryEMDaoTx := QEM.TX(tx)
	QEMList := make([]model.QueryExamMain, 0)
	var qem model.QueryExamMain
	for _, rowList := range datasAll {
		qem.ExamNum = rowList[0]
		qem.StudentName = rowList[1]
		qem.ClassName = rowList[2]
		qem.ChineseScore = rowList[3]
		qem.MathScore = rowList[4]
		qem.EnglishScore = rowList[5]
		qem.TotalScore = rowList[6]
		qem.ClassRate = rowList[7]
		qem.SchoolRate = rowList[8]
		qem.StepRank = rowList[9]
		qem.UploadDate = time.Now().Format("20060102")
		qem.IsDeleted = "0"
		QEMList = append(QEMList, qem)
	}

	sqlResult, err := QueryEMDaoTx.Insert(QEMList)
	if err != nil {
		logger.Errorf("插入数据库表[%s]数据失败[%s]", QEM.Table, err)
		tx.Rollback()
	}
	tx.Commit()
	if id, _ := sqlResult.RowsAffected(); id != 0 {
		logger.Infof("插入数据库表[%s]数据成功id[%d]", QEM.Table, id)
	}
	return nil
}

// FindTableOne 根据条件查询数据库中的某条数据，多个只返回第一条记录
func (QEM *queryExamMainDao) FindTableOne(logger *zap.SugaredLogger, where ...interface{}) (interface{}, error) {

	qem := (*model.QueryExamMain)(nil)
	err := QueryExamMain.Where(where).Struct(&qem)
	if err != nil {
		logger.Errorf("查询表[%s]语句%s，失败[%s]", QEM.Table, where, err)
	}
	return qem, nil
}

// FindTableAll 根据条件查询数据库中的某条数据，多个只返回第一条记录 qemResultList := qems.([]*model.QueryExamMain)
func (QEM *queryExamMainDao) FindTableAll(logger *zap.SugaredLogger, where ...interface{}) (interface{}, error) {

	qems := ([]*model.QueryExamMain)(nil)
	err := QueryExamMain.Where(where).Structs(&qems)
	if err != nil {
		logger.Errorf("查询表[%s]语句%s，失败[%s]", QEM.Table, where, err)
	}
	return qems, nil
}

// SoftDeletRecord 根据条件软删除数据库中多条数据(is_deleted=1,插入delete_time)
func (QEM *queryExamMainDao) SoftDeletRecord(logger *zap.SugaredLogger, where ...interface{}) error {
	db := g.DB()
	sqlResult, err := db.Table(QEM.Table).Data(where[0]).Data("delete_time", time.Now().Format("20060102150405")).Where(where[1:]).Update()
	if err != nil {
		logger.Errorf("软删除数据库表[%s]语句%s,数据失败[%s]", QEM.Table, where, err)
	}
	if id, _ := sqlResult.RowsAffected(); id != 0 {
		logger.Infof("软删除数据库表[%s]数据成功总数[%d]", QEM.Table, id)
	}

	return nil
}
