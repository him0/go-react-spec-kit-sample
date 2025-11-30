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
	db infrastructure.DBTX
}

// NewUserCommand UserCommandのコンストラクタ
func NewUserCommand(db infrastructure.DBTX) *UserCommand {
	return &UserCommand{db: db}
}

// WithTx トランザクション用の新しいUserCommandを返す
func (c *UserCommand) WithTx(tx infrastructure.DBTX) *UserCommand {
	return &UserCommand{db: tx}
}

// Create ユーザーを作成
func (c *UserCommand) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (id, name, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := c.db.ExecContext(
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

// Update ユーザーを更新
func (c *UserCommand) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := c.db.ExecContext(
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

// Delete ユーザーを削除
func (c *UserCommand) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := c.db.ExecContext(ctx, query, id)
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

// CreateWithTx トランザクション内でユーザーを作成
func CreateWithTx(ctx context.Context, tx infrastructure.DBTX, user *domain.User) error {
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

// UpdateWithTx トランザクション内でユーザーを更新
func UpdateWithTx(ctx context.Context, tx infrastructure.DBTX, user *domain.User) error {
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

// DeleteWithTx トランザクション内でユーザーを削除
func DeleteWithTx(ctx context.Context, tx infrastructure.DBTX, id string) error {
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

// Ensure *sql.DB implements DBTX
var _ infrastructure.DBTX = (*sql.DB)(nil)
