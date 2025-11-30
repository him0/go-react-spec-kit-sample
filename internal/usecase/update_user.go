package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
)

// UpdateUserUsecase ユーザー更新ユースケース
type UpdateUserUsecase struct {
	userCommand UserCommandRepository
	userQuery   UserQueryRepository
}

// NewUpdateUserUsecase UpdateUserUsecaseのコンストラクタ
func NewUpdateUserUsecase(
	userCommand UserCommandRepository,
	userQuery UserQueryRepository,
) *UpdateUserUsecase {
	return &UpdateUserUsecase{
		userCommand: userCommand,
		userQuery:   userQuery,
	}
}

// Execute ユーザーを更新
func (u *UpdateUserUsecase) Execute(ctx context.Context, id, name, email string) (*domain.User, error) {
	// 既存ユーザーの取得
	user, err := u.userQuery.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// メールアドレスが変更される場合、重複チェック
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

	// 永続化
	if err := u.userCommand.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
