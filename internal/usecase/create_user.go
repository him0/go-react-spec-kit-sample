package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/command"
	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
)

// CreateUserUsecase ユーザー作成ユースケース
type CreateUserUsecase struct {
	userQuery   UserQueryRepository
	userCommand TransactionManager
}

// NewCreateUserUsecase CreateUserUsecaseのコンストラクタ
func NewCreateUserUsecase(
	userQuery UserQueryRepository,
	userCommand TransactionManager,
) *CreateUserUsecase {
	return &CreateUserUsecase{
		userQuery:   userQuery,
		userCommand: userCommand,
	}
}

// Execute ユーザーを作成
func (u *CreateUserUsecase) Execute(ctx context.Context, name, email string) (*domain.User, error) {
	// メールアドレスの重複チェック（リーダーDB）
	existingUser, err := u.userQuery.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// ドメインモデルの作成
	user, err := domain.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	// 永続化（ライターDB、トランザクション内）
	err = u.userCommand.RunInTransaction(ctx, func(ctx context.Context, tx infrastructure.DBTX) error {
		return command.Create(ctx, tx, user)
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
