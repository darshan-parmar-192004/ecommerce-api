package unit

import (
	"context"
	"database/sql"
	"testing"

	"backend/internal/querybuilder"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestQueryBuilder_New(t *testing.T) {
	t.Run("creates QueryBuilder", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		qb := querybuilder.New(db, "products")
		assert.NotNil(t, qb)
	})
}

func TestQueryBuilder_Select(t *testing.T) {
	t.Run("selects specific columns", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id", "name", "price")

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})

	t.Run("select with no columns", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select()

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_Where(t *testing.T) {
	t.Run("adds where clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id").Where("product_id", "PROD-12345678")

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_OrderBy(t *testing.T) {
	t.Run("adds order by clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id").OrderBy("name")

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_OrderByDesc(t *testing.T) {
	t.Run("adds order by desc clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id").OrderByDesc("created_at")

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_Limit(t *testing.T) {
	t.Run("adds limit clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id").Limit(10)

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_Offset(t *testing.T) {
	t.Run("adds offset clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id").Offset(20)

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

func TestQueryBuilder_QueryRow(t *testing.T) {
	t.Run("returns single row", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"product_id", "name"}).AddRow("PROD-12345678", "Test"))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id", "name").Where("product_id", "PROD-12345678")

		row := qb.QueryRow(context.Background())
		assert.NotNil(t, row)
	})
}

func TestQueryBuilder_Count(t *testing.T) {
	t.Run("returns count", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(5))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id")

		count, err := qb.Count(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 5, count)
	})
}

func TestQueryBuilder_Exec(t *testing.T) {
	t.Run("executes query", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("SELECT").WillReturnResult(sqlmock.NewResult(0, 0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id")

		result, err := qb.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestQueryBuilder_Chaining(t *testing.T) {
	t.Run("chains multiple methods", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		qb := querybuilder.New(db, "products")
		qb.Select("product_id", "name", "price").
			Where("category_id", "CAT-12345678").
			OrderBy("name").
			Limit(10).
			Offset(0)

		rows, err := qb.Query(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, rows)
		_ = rows.Close()
	})
}

// ─── InsertBuilder Tests ─────────────────────────────────────────────────────

func TestInsertBuilder_New(t *testing.T) {
	t.Run("creates InsertBuilder", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		ib := querybuilder.NewInsert(db, "products")
		assert.NotNil(t, ib)
	})
}

func TestInsertBuilder_Columns(t *testing.T) {
	t.Run("sets columns", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))

		ib := querybuilder.NewInsert(db, "products")
		ib.Columns("product_id", "name", "price")

		result, err := ib.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInsertBuilder_Values(t *testing.T) {
	t.Run("sets values", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))

		ib := querybuilder.NewInsert(db, "products")
		ib.Columns("product_id", "name", "price").
			Values("PROD-12345678", "Test Product", 9.99)

		result, err := ib.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInsertBuilder_Record(t *testing.T) {
	t.Run("sets record", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))

		ib := querybuilder.NewInsert(db, "products")
		ib.Record(map[string]interface{}{
			"product_id": "PROD-12345678",
			"name":       "Test Product",
			"price":      9.99,
		})

		result, err := ib.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestInsertBuilder_Exec(t *testing.T) {
	t.Run("executes insert", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))

		ib := querybuilder.NewInsert(db, "products")
		ib.Columns("product_id", "name").
			Values("PROD-12345678", "Test")

		result, err := ib.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

// ─── UpdateBuilder Tests ─────────────────────────────────────────────────────

func TestUpdateBuilder_New(t *testing.T) {
	t.Run("creates UpdateBuilder", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		ub := querybuilder.NewUpdate(db, "products")
		assert.NotNil(t, ub)
	})
}

func TestUpdateBuilder_Set(t *testing.T) {
	t.Run("sets column value", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))

		ub := querybuilder.NewUpdate(db, "products")
		ub.Set("name", "Updated Product")

		result, err := ub.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdateBuilder_Where(t *testing.T) {
	t.Run("adds where clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))

		ub := querybuilder.NewUpdate(db, "products")
		ub.Set("name", "Updated").Where("product_id", "PROD-12345678")

		result, err := ub.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestUpdateBuilder_Exec(t *testing.T) {
	t.Run("executes update", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(0, 1))

		ub := querybuilder.NewUpdate(db, "products")
		ub.Set("name", "Updated").
			Set("price", 19.99).
			Where("product_id", "PROD-12345678")

		result, err := ub.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

// ─── DeleteBuilder Tests ─────────────────────────────────────────────────────

func TestDeleteBuilder_New(t *testing.T) {
	t.Run("creates DeleteBuilder", func(t *testing.T) {
		db, _, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		db2 := querybuilder.NewDelete(db, "products")
		assert.NotNil(t, db2)
	})
}

func TestDeleteBuilder_Where(t *testing.T) {
	t.Run("adds where clause", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(0, 1))

		db2 := querybuilder.NewDelete(db, "products")
		db2.Where("product_id", "PROD-12345678")

		result, err := db2.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

func TestDeleteBuilder_Exec(t *testing.T) {
	t.Run("executes delete", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		defer func() { _ = db.Close() }()

		mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(0, 1))

		db2 := querybuilder.NewDelete(db, "products")
		db2.Where("product_id", "PROD-12345678")

		result, err := db2.Exec(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

// Suppress unused import warning
var _ = sql.ErrNoRows
