package usecase

import (
	"context"
	"errors"
)

// DeleteUserUsecase ユーザー削除ユースケース
type DeleteUserUsecase struct {
	userCommand UserCommandRepository
	userQuery   UserQueryRepository
}

// NewDeleteUserUsecase DeleteUserUsecaseのコンストラクタ
func NewDeleteUserUsecase(
	userCommand UserCommandRepository,
	userQuery UserQueryRepository,
) *DeleteUserUsecase {
	return &DeleteUserUsecase{
		userCommand: userCommand,
		userQuery:   userQuery,
	}
}

// Execute ユーザーを削除
func (u *DeleteUserUsecase) Execute(ctx context.Context, id string) error {
	// 存在確認
	user, err := u.userQuery.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// 削除
	return u.userCommand.Delete(ctx, id)
}
