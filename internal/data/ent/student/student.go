// Code generated by entc, DO NOT EDIT.

package student

const (
	// Label holds the string label denoting the student type in the database.
	Label = "student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldExamNum holds the string denoting the exam_num field in the database.
	FieldExamNum = "exam_num"
	// FieldStudentName holds the string denoting the student_name field in the database.
	FieldStudentName = "student_name"
	// FieldClassName holds the string denoting the class_name field in the database.
	FieldClassName = "class_name"
	// FieldChineseScore holds the string denoting the chinese_score field in the database.
	FieldChineseScore = "chinese_score"
	// FieldMathScore holds the string denoting the math_score field in the database.
	FieldMathScore = "math_score"
	// FieldEnglishScore holds the string denoting the english_score field in the database.
	FieldEnglishScore = "english_score"
	// FieldTotalScore holds the string denoting the total_score field in the database.
	FieldTotalScore = "total_score"
	// FieldClassRate holds the string denoting the class_rate field in the database.
	FieldClassRate = "class_rate"
	// FieldSchoolRate holds the string denoting the school_rate field in the database.
	FieldSchoolRate = "school_rate"
	// FieldStepRank holds the string denoting the step_rank field in the database.
	FieldStepRank = "step_rank"
	// FieldUploadDate holds the string denoting the upload_date field in the database.
	FieldUploadDate = "upload_date"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldDeleteTime holds the string denoting the delete_time field in the database.
	FieldDeleteTime = "delete_time"

	// Table holds the table name of the student in the database.
	Table = "query_exam_main"
)

// Columns holds all SQL columns for student fields.
var Columns = []string{
	FieldID,
	FieldExamNum,
	FieldStudentName,
	FieldClassName,
	FieldChineseScore,
	FieldMathScore,
	FieldEnglishScore,
	FieldTotalScore,
	FieldClassRate,
	FieldSchoolRate,
	FieldStepRank,
	FieldUploadDate,
	FieldIsDeleted,
	FieldDeleteTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ExamNumValidator is a validator for the "exam_num" field. It is called by the builders before save.
	ExamNumValidator func(string) error
	// DefaultChineseScore holds the default value on creation for the "chinese_score" field.
	DefaultChineseScore string
	// DefaultMathScore holds the default value on creation for the "math_score" field.
	DefaultMathScore string
	// DefaultEnglishScore holds the default value on creation for the "english_score" field.
	DefaultEnglishScore string
	// DefaultTotalScore holds the default value on creation for the "total_score" field.
	DefaultTotalScore string
	// DefaultClassRate holds the default value on creation for the "class_rate" field.
	DefaultClassRate string
	// DefaultSchoolRate holds the default value on creation for the "school_rate" field.
	DefaultSchoolRate string
	// DefaultStepRank holds the default value on creation for the "step_rank" field.
	DefaultStepRank string
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted string
	// DefaultDeleteTime holds the default value on creation for the "delete_time" field.
	DefaultDeleteTime string
)