package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
)

// CreateUserUsecase ユーザー作成ユースケース
type CreateUserUsecase struct {
	userCommand UserCommandRepository
	userQuery   UserQueryRepository
}

// NewCreateUserUsecase CreateUserUsecaseのコンストラクタ
func NewCreateUserUsecase(
	userCommand UserCommandRepository,
	userQuery UserQueryRepository,
) *CreateUserUsecase {
	return &CreateUserUsecase{
		userCommand: userCommand,
		userQuery:   userQuery,
	}
}

// Execute ユーザーを作成
func (u *CreateUserUsecase) Execute(ctx context.Context, name, email string) (*domain.User, error) {
	// メールアドレスの重複チェック
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

	// 永続化
	if err := u.userCommand.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
