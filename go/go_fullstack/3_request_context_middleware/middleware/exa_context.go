package middleware

import (
	"context"
	"net/http"
)

// 从请求中获取内容, 设置到请求上下文中
func newExampleContext(ctx context.Context, r *http.Request) context.Context {
	fooId := r.Header.Get("X-Foo-Id")
	if fooId == "" {
		fooId = "bar"
	}
	// 设置值
	return context.WithValue(ctx, "fooId", fooId)
}

// 接收处理器,返回处理器
func ContextExapmleHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// 调用newExampleContext 设置值
			ctx := newExampleContext(r.Context(), r)
			next.ServeHTTP(w, r.WithContext(ctx))
		}()
	})
}
