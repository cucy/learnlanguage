package helpers

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware wraps the logs for incoming requests 日志中间件为传入的请求包装日志
type LoggingMiddleware struct {
	Logger log.Logger
	Next   EncryptService
}

// Encrypt logs the encyption requests 加密请求的日志
func (mw LoggingMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Encrypt(ctx, key, text)
	return
}

// Decrypt logs the encyption requests
func (mw LoggingMiddleware) Decrypt(ctx context.Context, key string, text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"message", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
