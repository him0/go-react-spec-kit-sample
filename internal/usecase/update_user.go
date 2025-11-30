package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/command"
	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
)

// UpdateUserUsecase ユーザー更新ユースケース
type UpdateUserUsecase struct {
	userQuery   UserQueryRepository
	userCommand TransactionManager
}

// NewUpdateUserUsecase UpdateUserUsecaseのコンストラクタ
func NewUpdateUserUsecase(
	userQuery UserQueryRepository,
	userCommand TransactionManager,
) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		userQuery:   userQuery,
		userCommand: userCommand,
	}
}

// Execute ユーザーを更新
func (u *UpdateUserUsecase) Execute(ctx context.Context, id, name, email string) (*domain.User, error) {
	// 既存ユーザーの取得（リーダーDB）
	user, err := u.userQuery.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// メールアドレスが変更される場合、重複チェック（リーダーDB）
	if email != "" && email != user.Email {
		existingUser, err := u.userQuery.FindByEmail(ctx, email)
		if err != nil {
			return nil, err
		}
		if existingUser != nil {
			return nil, errors.New("email already exists")
		}
	}

	// ドメインモデルの更新
	if err := user.Update(name, email); err != nil {
		return nil, err
	}

	// 永続化（ライターDB、トランザクション内）
	err = u.userCommand.RunInTransaction(ctx, func(ctx context.Context, tx infrastructure.DBTX) error {
		return command.Update(ctx, tx, user)
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
