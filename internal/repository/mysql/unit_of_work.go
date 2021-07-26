package mysql

import (
	"context"
	"database/sql"
	"errors"
)

type MySQLTransactionKey string

type MySQLUnitOfWork struct {
	DB *sql.DB
}

func NewMySQLUnitOfWork(DB *sql.DB) *MySQLUnitOfWork {
	return &MySQLUnitOfWork{DB}
}

func (muow MySQLUnitOfWork) Begin(ctx context.Context) (context.Context, error) {
	tx, err := muow.DB.BeginTx(ctx, nil)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, MySQLTransactionKey("tx"), tx), nil
}

func (muow MySQLUnitOfWork) Commit(ctx context.Context) error {
	tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx)
	if !ok {
		return errors.New("failed get transcation context")
	}

	return tx.Commit()
}

func (muow MySQLUnitOfWork) Rollback(ctx context.Context) error {
	tx, ok := ctx.Value(MySQLTransactionKey("tx")).(*sql.Tx)
	if !ok {
		return errors.New("failed get transcation context")
	}

	return tx.Rollback()
}
