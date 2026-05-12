package querybuilder

import (
	"context"
	"database/sql"

	"gopkg.in/doug-martin/goqu.v5"
)

type QueryBuilder struct {
	db *sql.DB
}

var qb *QueryBuilder

func New(db *sql.DB) *QueryBuilder {
	if qb == nil {
		qb = &QueryBuilder{
			db: db,
		}
	}
	return qb
}

func From(table string) *goqu.Dataset {
	return goqu.From(table)
}

func FromAliased(table goqu.AliasedExpression) *goqu.Dataset {
	return goqu.From(table)
}

func Ex(m map[string]interface{}) goqu.Ex {
	return goqu.Ex(m)
}

func I(ident string) goqu.IdentifierExpression {
	return goqu.I(ident)
}

var SUM = goqu.SUM
var COALESCE = goqu.COALESCE
var COUNT = goqu.COUNT
var L = goqu.L
var On = goqu.On

func As(expr goqu.IdentifierExpression, alias string) interface{} {
	return expr.As(alias)
}

func Record(m map[string]interface{}) map[string]interface{} {
	return m
}

func (q *QueryBuilder) Exec(ctx context.Context, sqlStr string, args ...interface{}) (sql.Result, error) {
	return q.db.ExecContext(ctx, sqlStr, args...)
}

func (q *QueryBuilder) QueryContext(ctx context.Context, sqlStr string, args ...interface{}) (*sql.Rows, error) {
	return q.db.QueryContext(ctx, sqlStr, args...)
}

func (q *QueryBuilder) QueryRowContext(ctx context.Context, sqlStr string, args ...interface{}) *sql.Row {
	return q.db.QueryRowContext(ctx, sqlStr, args...)
}

func ToSQL(ds *goqu.Dataset) (string, []interface{}) {
	sqlStr, args, err := ds.ToSql()
	if err != nil {
		panic("failed to generate SQL: " + err.Error())
	}
	return sqlStr, args
}

func GetSQL(crud *goqu.CrudExec) (string, []interface{}) {
	return crud.Sql, crud.Args
}

func NewOp(m map[string]interface{}) goqu.Op {
	return goqu.Op(m)
}
