package usecase

import (
	"context"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/infrastructure"
)

// TransactionManager トランザクション管理のインターフェース
type TransactionManager interface {
	RunInTransaction(ctx context.Context, fn func(ctx context.Context, tx infrastructure.DBTX) error) error
}

// UserQueryRepository 読み取り操作のインターフェース
type UserQueryRepository interface {
	FindByID(ctx context.Context, id string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindAll(ctx context.Context, limit, offset int) ([]*domain.User, error)
	Count(ctx context.Context) (int, error)
}
