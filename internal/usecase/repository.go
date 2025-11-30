package usecase

import (
	"context"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
)

// UserCommandRepository 書き込み操作のインターフェース
type UserCommandRepository interface {
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

// UserQueryRepository 読み取り操作のインターフェース
type UserQueryRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindAll(ctx context.Context, limit, offset int) ([]*domain.User, error)
	Count(ctx context.Context) (int, error)
}
