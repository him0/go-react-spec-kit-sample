package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/command"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
	"github.com/example/go-react-spec-kit-sample/internal/queryservice"
)

// DeleteUserUsecase ユーザー削除ユースケース
type DeleteUserUsecase struct {
	txManager TransactionManager
}

// NewDeleteUserUsecase DeleteUserUsecaseのコンストラクタ
func NewDeleteUserUsecase(txManager TransactionManager) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		txManager: txManager,
	}
}

// Execute ユーザーを削除
func (u *DeleteUserUsecase) Execute(ctx context.Context, id string) error {
	return u.txManager.RunInTransaction(ctx, func(ctx context.Context, tx infrastructure.DBTX) error {
		// 存在確認
		user, err := queryservice.FindByIDWithTx(ctx, tx, id)
		if err != nil {
			return err
		}
		if user == nil {
			return errors.New("user not found")
		}

		// 削除
		return command.DeleteWithTx(ctx, tx, id)
	})
}
