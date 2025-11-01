//go:build tools
// +build tools

// このファイルは開発時に使用するツールの依存関係を宣言します。
// go.modにツールの依存関係を記録するために使用されます。

package tools

import (
	_ "github.com/sqlc-dev/sqlc/cmd/sqlc"
)
