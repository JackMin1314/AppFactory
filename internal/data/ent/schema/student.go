package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Annotations of the User.
func (Student) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "query_exam_main"},
	}
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("exam_num").
			NotEmpty().
			Unique(),
		field.String("student_name"),
		field.String("class_name"),
		field.String("chinese_score").
			Default("0"),
		field.String("math_score").
			Default("0"),
		field.String("english_score").
			Default("0"),
		field.String("total_score").
			Default("0"),
		field.String("class_rate").
			Default(""),
		field.String("school_rate").
			Default(""),
		field.String("step_rank").
			Default(""),
		field.String("upload_date"),
		field.String("is_deleted").
			Default("0"),
		field.String("delete_time").
			Default(""),
		// field.Time("created_at").
		// 	Default(time.Now).SchemaType(map[string]string{
		// 	dialect.MySQL: "datetime",
		// }),
		// field.Time("updated_at").
		// 	Default(time.Now).SchemaType(map[string]string{
		// 	dialect.MySQL: "datetime",
		// }),
	}
}

// Edges of the Student.
func (Student) Edges() []ent.Edge {
	return nil
}
