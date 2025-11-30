package command

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
)

// UserCommand ユーザー書き込み操作を担当
type UserCommand struct {
	db *sql.DB
}

// NewUserCommand UserCommandのコンストラクタ
func NewUserCommand(db *sql.DB) *UserCommand {
	return &UserCommand{db: db}
}

// RunInTransaction トランザクション内で処理を実行
func (c *UserCommand) RunInTransaction(ctx context.Context, fn func(ctx context.Context, tx infrastructure.DBTX) error) error {
	tx, err := c.db.BeginTx(ctx, nil)
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

// Create ユーザーを作成（トランザクション内で使用）
func Create(ctx context.Context, tx infrastructure.DBTX, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		user.ID,
		user.Name,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

// Update ユーザーを更新（トランザクション内で使用）
func Update(ctx context.Context, tx infrastructure.DBTX, user *domain.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := tx.ExecContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found: %s", user.ID)
	}

	return nil
}

// Delete ユーザーを削除（トランザクション内で使用）
func Delete(ctx context.Context, tx infrastructure.DBTX, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found: %s", id)
	}

	return nil
}
