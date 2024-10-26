package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type FakeDialector struct {
	dsn string
}

func Open(dsn string) gorm.Dialector {
	return &FakeDialector{
		dsn: dsn,
	}
}

func (dialector *FakeDialector) Name() string {
	return "fake"
}

func (dialector *FakeDialector) Initialize(db *gorm.DB) error {
	callbacks.RegisterDefaultCallbacks(db, &callbacks.Config{
		CreateClauses:        []string{"INSERT", "VALUES", "ON CONFLICT", "RETURNING"},
		UpdateClauses:        []string{"UPDATE", "SET", "WHERE", "RETURNING"},
		DeleteClauses:        []string{"DELETE", "FROM", "WHERE", "RETURNING"},
		LastInsertIDReversed: true,
	})

	return nil
}

func (dialector *FakeDialector) DefaultValueOf(field *schema.Field) clause.Expression {
	return clause.Expr{SQL: "DEFAULT"}
}

func (dialector *FakeDialector) Migrator(*gorm.DB) gorm.Migrator {
	return nil
}

func (dialector *FakeDialector) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	writer.WriteByte('?') //nolint:errcheck
}

func (dialector *FakeDialector) QuoteTo(writer clause.Writer, str string) {
	var (
		underQuoted, selfQuoted bool
		continuousBacktick      int8
		shiftDelimiter          int8
	)

	for _, v := range []byte(str) {
		switch v {
		case '`':
			continuousBacktick++
			if continuousBacktick == 2 {
				writer.WriteString("``") //nolint:errcheck
				continuousBacktick = 0
			}
		case '.':
			if continuousBacktick > 0 || !selfQuoted {
				shiftDelimiter = 0
				underQuoted = false
				continuousBacktick = 0
				writer.WriteByte('`') //nolint:errcheck
			}
			writer.WriteByte(v) //nolint:errcheck
			continue
		default:
			if shiftDelimiter-continuousBacktick <= 0 && !underQuoted {
				writer.WriteByte('`') //nolint:errcheck
				underQuoted = true
				if selfQuoted = continuousBacktick > 0; selfQuoted {
					continuousBacktick -= 1
				}
			}

			for ; continuousBacktick > 0; continuousBacktick -= 1 {
				writer.WriteString("``") //nolint:errcheck
			}

			writer.WriteByte(v) //nolint:errcheck
		}
		shiftDelimiter++
	}

	if continuousBacktick > 0 && !selfQuoted {
		writer.WriteString("``") //nolint:errcheck
	}
	writer.WriteByte('`') //nolint:errcheck
}

func (dialector *FakeDialector) Explain(sql string, vars ...interface{}) string {
	return logger.ExplainSQL(sql, nil, `"`, vars...)
}

func (dialector *FakeDialector) DataTypeOf(*schema.Field) string {
	return ""
}

func (dialector *FakeDialector) Translate(err error) error {
	return err
}
