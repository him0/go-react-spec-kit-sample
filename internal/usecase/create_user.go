package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/command"
	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
	"github.com/example/go-react-spec-kit-sample/internal/queryservice"
)

// CreateUserUsecase ユーザー作成ユースケース
type CreateUserUsecase struct {
	txManager TransactionManager
}

// NewCreateUserUsecase CreateUserUsecaseのコンストラクタ
func NewCreateUserUsecase(txManager TransactionManager) *CreateUserUsecase {
	return &CreateUserUsecase{
		txManager: txManager,
	}
}

// Execute ユーザーを作成
func (u *CreateUserUsecase) Execute(ctx context.Context, name, email string) (*domain.User, error) {
	var createdUser *domain.User

	err := u.txManager.RunInTransaction(ctx, func(ctx context.Context, tx infrastructure.DBTX) error {
		// メールアドレスの重複チェック
		existingUser, err := queryservice.FindByEmailWithTx(ctx, tx, email)
		if err != nil {
			return err
		}
		if existingUser != nil {
			return errors.New("email already exists")
		}

		// ドメインモデルの作成
		user, err := domain.NewUser(name, email)
		if err != nil {
			return err
		}

		// 永続化
		if err := command.CreateWithTx(ctx, tx, user); err != nil {
			return err
		}

		createdUser = user
		return nil
	})

	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
