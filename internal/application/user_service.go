package application

import (
	"errors"
	"log/slog"

	"github.com/example/go-react-spec-kit-sample/internal/domain"
	"github.com/example/go-react-spec-kit-sample/internal/pkg/logger"
)

// UserService アプリケーションサービス
type UserService struct {
	userRepo domain.UserRepository
	logger   *slog.Logger
}

// NewUserService UserServiceのコンストラクタ
func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		logger:   logger.Get(),
	}
}

// CreateUser ユーザーを作成
func (s *UserService) CreateUser(name, email string) (*domain.User, error) {
	s.logger.Debug("service: creating user",
		slog.String("name", name),
		slog.String("email", email),
	)

	user, err := domain.NewUser(name, email)
	if err != nil {
		s.logger.Warn("service: invalid user data",
			slog.String("error", err.Error()),
			slog.String("name", name),
			slog.String("email", email),
		)
		return nil, err
	}

	if err := s.userRepo.Save(user); err != nil {
		s.logger.Error("service: failed to save user",
			slog.String("error", err.Error()),
			slog.String("user_id", user.ID),
		)
		return nil, err
	}

	s.logger.Info("service: user created",
		slog.String("user_id", user.ID),
		slog.String("email", user.Email),
	)

	return user, nil
}

// GetUser ユーザーを取得
func (s *UserService) GetUser(id string) (*domain.User, error) {
	s.logger.Debug("service: getting user", slog.String("user_id", id))
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		s.logger.Error("service: failed to get user",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
	}
	return user, err
}

// ListUsers ユーザー一覧を取得
func (s *UserService) ListUsers(limit, offset int) ([]*domain.User, int, error) {
	s.logger.Debug("service: listing users",
		slog.Int("limit", limit),
		slog.Int("offset", offset),
	)

	users, err := s.userRepo.FindAll(limit, offset)
	if err != nil {
		s.logger.Error("service: failed to list users",
			slog.String("error", err.Error()),
		)
		return nil, 0, err
	}

	total, err := s.userRepo.Count()
	if err != nil {
		s.logger.Error("service: failed to count users",
			slog.String("error", err.Error()),
		)
		return nil, 0, err
	}

	return users, total, nil
}

// UpdateUser ユーザーを更新
func (s *UserService) UpdateUser(id, name, email string) (*domain.User, error) {
	s.logger.Debug("service: updating user",
		slog.String("user_id", id),
		slog.String("name", name),
		slog.String("email", email),
	)

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		s.logger.Error("service: failed to find user for update",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
		return nil, err
	}

	if user == nil {
		s.logger.Warn("service: user not found for update",
			slog.String("user_id", id),
		)
		return nil, errors.New("user not found")
	}

	if err := user.Update(name, email); err != nil {
		s.logger.Warn("service: invalid update data",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
		return nil, err
	}

	if err := s.userRepo.Save(user); err != nil {
		s.logger.Error("service: failed to save updated user",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
		return nil, err
	}

	s.logger.Info("service: user updated",
		slog.String("user_id", id),
		slog.String("email", user.Email),
	)

	return user, nil
}

// DeleteUser ユーザーを削除
func (s *UserService) DeleteUser(id string) error {
	s.logger.Debug("service: deleting user", slog.String("user_id", id))

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		s.logger.Error("service: failed to find user for deletion",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
		return err
	}

	if user == nil {
		s.logger.Warn("service: user not found for deletion",
			slog.String("user_id", id),
		)
		return errors.New("user not found")
	}

	if err := s.userRepo.Delete(id); err != nil {
		s.logger.Error("service: failed to delete user",
			slog.String("error", err.Error()),
			slog.String("user_id", id),
		)
		return err
	}

	s.logger.Info("service: user deleted", slog.String("user_id", id))

	return nil
}
