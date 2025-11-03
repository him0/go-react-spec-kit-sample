package infrastructure

import (
	"errors"
	"log/slog"
	"sync"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/pkg/logger"
)

// InMemoryUserRepository インメモリユーザーリポジトリ
type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*domain.User
	logger *slog.Logger
}

// NewInMemoryUserRepository InMemoryUserRepositoryのコンストラクタ
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[string]*domain.User),
		logger: logger.Get(),
	}
}

// FindByID IDでユーザーを検索
func (r *InMemoryUserRepository) FindByID(id string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug("repository: finding user by id", slog.String("user_id", id))

	user, ok := r.users[id]
	if !ok {
		r.logger.Debug("repository: user not found", slog.String("user_id", id))
		return nil, errors.New("user not found")
	}

	return user, nil
}

// FindAll 全ユーザーを取得
func (r *InMemoryUserRepository) FindAll(limit, offset int) ([]*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.logger.Debug("repository: finding all users",
		slog.Int("limit", limit),
		slog.Int("offset", offset),
	)

	users := make([]*domain.User, 0)
	i := 0
	for _, user := range r.users {
		if i >= offset && i < offset+limit {
			users = append(users, user)
		}
		i++
		if i >= offset+limit {
			break
		}
	}

	r.logger.Debug("repository: found users", slog.Int("count", len(users)))

	return users, nil
}

// Count ユーザー数を取得
func (r *InMemoryUserRepository) Count() (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	count := len(r.users)
	r.logger.Debug("repository: counting users", slog.Int("total", count))

	return count, nil
}

// Save ユーザーを保存
func (r *InMemoryUserRepository) Save(user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug("repository: saving user",
		slog.String("user_id", user.ID),
		slog.String("email", user.Email),
	)

	r.users[user.ID] = user

	r.logger.Info("repository: user saved", slog.String("user_id", user.ID))

	return nil
}

// Delete ユーザーを削除
func (r *InMemoryUserRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.logger.Debug("repository: deleting user", slog.String("user_id", id))

	if _, ok := r.users[id]; !ok {
		r.logger.Warn("repository: user not found for deletion", slog.String("user_id", id))
		return errors.New("user not found")
	}

	delete(r.users, id)

	r.logger.Info("repository: user deleted", slog.String("user_id", id))

	return nil
}
