package logger

import (
	"log/slog"
	"os"
	"strings"
)

// グローバルロガー
var defaultLogger *slog.Logger

// Setup はロガーを初期化します
// LOG_LEVEL環境変数でログレベルを制御できます（DEBUG, INFO, WARN, ERROR）
// LOG_FORMAT環境変数でフォーマットを制御できます（json, text）デフォルトはjson
func Setup() *slog.Logger {
	// ログレベルの設定
	level := getLogLevel()

	// ログフォーマットの設定
	format := os.Getenv("LOG_FORMAT")
	if format == "" {
		format = "json"
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
		// ソースコードの位置情報を追加（開発時に便利）
		AddSource: level == slog.LevelDebug,
	}

	if format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	defaultLogger = logger
	slog.SetDefault(logger)

	return logger
}

// Get はグローバルロガーを返します
func Get() *slog.Logger {
	if defaultLogger == nil {
		return Setup()
	}
	return defaultLogger
}

// getLogLevel は環境変数からログレベルを取得します
func getLogLevel() slog.Level {
	levelStr := strings.ToUpper(os.Getenv("LOG_LEVEL"))
	switch levelStr {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
