package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
)

// DBTX は *sql.DB と *sql.Tx の共通インターフェース
type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// TransactionManager トランザクション管理
type TransactionManager struct {
	db *sql.DB
}

// NewTransactionManager TransactionManagerのコンストラクタ
func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

// RunInTransaction トランザクション内で処理を実行
func (tm *TransactionManager) RunInTransaction(ctx context.Context, fn func(ctx context.Context, tx DBTX) error) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	if err := fn(ctx, tx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("failed to rollback: %v (original error: %w)", rbErr, err)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
