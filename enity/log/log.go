// Package log 用于隔离核心代码与外部设施的依赖
package log

import (
	"context"
)

func Errorf(format string, args ...interface{}) {
}

func ErrorContextf(ctx context.Context, format string, args ...interface{}) {
}
