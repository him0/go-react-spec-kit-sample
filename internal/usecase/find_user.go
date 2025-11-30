package usecase

import (
	"context"
	"errors"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
)

// FindUserUsecase ユーザー取得ユースケース
type FindUserUsecase struct {
	userQuery UserQueryRepository
}

// NewFindUserUsecase FindUserUsecaseのコンストラクタ
func NewFindUserUsecase(userQuery UserQueryRepository) *FindUserUsecase {
	return &FindUserUsecase{
		userQuery: userQuery,
	}
}

// Execute ユーザーを取得
func (u *FindUserUsecase) Execute(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userQuery.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
