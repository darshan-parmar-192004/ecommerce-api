package querybuilder

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type QueryBuilder struct {
	db      *sql.DB
	selects *goqu.SelectDataset
}

func New(db *sql.DB, table string) *QueryBuilder {
	return &QueryBuilder{
		db:      db,
		selects: goqu.Dialect("postgres").From(table),
	}
}

func (q *QueryBuilder) Select(cols ...string) *QueryBuilder {
	if len(cols) > 0 {
		colsInterface := make([]interface{}, len(cols))
		for i, c := range cols {
			colsInterface[i] = c
		}
		q.selects = q.selects.Select(colsInterface...)
	}
	return q
}

func (q *QueryBuilder) Where(field string, value interface{}) *QueryBuilder {
	q.selects = q.selects.Where(goqu.C(field).Eq(value))
	return q
}

func (q *QueryBuilder) OrderBy(field string) *QueryBuilder {
	q.selects = q.selects.Order(goqu.C(field).Asc())
	return q
}

func (q *QueryBuilder) OrderByDesc(field string) *QueryBuilder {
	q.selects = q.selects.Order(goqu.C(field).Desc())
	return q
}

func (q *QueryBuilder) Limit(limit int) *QueryBuilder {
	q.selects = q.selects.Limit(uint(limit))
	return q
}

func (q *QueryBuilder) Offset(offset int) *QueryBuilder {
	q.selects = q.selects.Offset(uint(offset))
	return q
}

func (q *QueryBuilder) Query(ctx context.Context) (*sql.Rows, error) {
	sqlStr, args, err := q.selects.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	return q.db.QueryContext(ctx, sqlStr, args...)
}

func (q *QueryBuilder) QueryRow(ctx context.Context) *sql.Row {
	sqlStr, args, _ := q.selects.Prepared(true).ToSQL()
	return q.db.QueryRowContext(ctx, sqlStr, args...)
}

func (q *QueryBuilder) Exec(ctx context.Context) (sql.Result, error) {
	sqlStr, args, err := q.selects.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	return q.db.ExecContext(ctx, sqlStr, args...)
}

func (q *QueryBuilder) Count(ctx context.Context) (int, error) {
	var count int
	sqlStr, args, _ := q.selects.Prepared(true).Select(goqu.COUNT(goqu.C("*"))).ToSQL()
	err := q.db.QueryRowContext(ctx, sqlStr, args...).Scan(&count)
	return count, err
}

type InsertBuilder struct {
	db      *sql.DB
	inserts *goqu.InsertDataset
}

func NewInsert(db *sql.DB, table string) *InsertBuilder {
	return &InsertBuilder{
		db:      db,
		inserts: goqu.Dialect("postgres").Insert(table),
	}
}

func (i *InsertBuilder) Columns(cols ...string) *InsertBuilder {
	colsInterface := make([]interface{}, len(cols))
	for idx, c := range cols {
		colsInterface[idx] = c
	}
	i.inserts = i.inserts.Cols(colsInterface...)
	return i
}

func (i *InsertBuilder) Values(vals ...interface{}) *InsertBuilder {
	vals2D := [][]interface{}{vals}
	i.inserts = i.inserts.Vals(vals2D...)
	return i
}

func (i *InsertBuilder) Record(record map[string]interface{}) *InsertBuilder {
	i.inserts = i.inserts.Rows(goqu.Record(record))
	return i
}

func (i *InsertBuilder) Exec(ctx context.Context) (sql.Result, error) {
	sqlStr, args, err := i.inserts.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	return i.db.ExecContext(ctx, sqlStr, args...)
}

type UpdateBuilder struct {
	db      *sql.DB
	updates *goqu.UpdateDataset
}

func NewUpdate(db *sql.DB, table string) *UpdateBuilder {
	return &UpdateBuilder{
		db:      db,
		updates: goqu.Dialect("postgres").Update(table),
	}
}

func (u *UpdateBuilder) Set(col string, val interface{}) *UpdateBuilder {
	u.updates = u.updates.Set(goqu.Record{col: val})
	return u
}

func (u *UpdateBuilder) Where(field string, value interface{}) *UpdateBuilder {
	u.updates = u.updates.Where(goqu.C(field).Eq(value))
	return u
}

func (u *UpdateBuilder) Exec(ctx context.Context) (sql.Result, error) {
	sqlStr, args, err := u.updates.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	return u.db.ExecContext(ctx, sqlStr, args...)
}

type DeleteBuilder struct {
	db      *sql.DB
	deletes *goqu.DeleteDataset
}

func NewDelete(db *sql.DB, table string) *DeleteBuilder {
	return &DeleteBuilder{
		db:      db,
		deletes: goqu.Dialect("postgres").Delete(table),
	}
}

func (d *DeleteBuilder) Where(field string, value interface{}) *DeleteBuilder {
	d.deletes = d.deletes.Where(goqu.C(field).Eq(value))
	return d
}

func (d *DeleteBuilder) Exec(ctx context.Context) (sql.Result, error) {
	sqlStr, args, err := d.deletes.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}
	return d.db.ExecContext(ctx, sqlStr, args...)
}
