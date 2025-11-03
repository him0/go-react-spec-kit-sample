package logger

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDKey contextKey = "request_id"
	loggerKey    contextKey = "logger"
)

// WithRequestID はコンテキストにリクエストIDを追加します
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// GetRequestID はコンテキストからリクエストIDを取得します
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}

// GenerateRequestID は新しいリクエストIDを生成します
func GenerateRequestID() string {
	return uuid.New().String()
}

// WithLogger はコンテキストにロガーを追加します
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

// FromContext はコンテキストからロガーを取得します
// コンテキストにリクエストIDがある場合は、自動的にログに追加します
func FromContext(ctx context.Context) *slog.Logger {
	// コンテキストからロガーを取得
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return enrichLogger(ctx, logger)
	}

	// デフォルトロガーを使用
	return enrichLogger(ctx, Get())
}

// enrichLogger はロガーにコンテキスト情報を追加します
func enrichLogger(ctx context.Context, logger *slog.Logger) *slog.Logger {
	requestID := GetRequestID(ctx)
	if requestID != "" {
		return logger.With(slog.String("request_id", requestID))
	}
	return logger
}
