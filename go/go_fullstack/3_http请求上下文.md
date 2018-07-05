

```go
// main.go

http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))
```

# 设置
```go
package middleware

import (
	"context"
	"net/http"
)

func newExampleContext(ctx context.Context, r *http.Request) context.Context {

	fooID := r.Header.Get("X-Foo-ID")
	if fooID == "" {
		fooID = "bar"
	}

	return context.WithValue(ctx, "fooID", fooID)

}

func ContextExampleHandler(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
             // 设置值
			ctx := newExampleContext(r.Context(), r)
			next.ServeHTTP(w, r.WithContext(ctx))

		}()

	})

}
```


# 获取

```go
package handlers

import (
	"net/http"
)

func FooHandler(w http.ResponseWriter, r *http.Request) {

	fooID := r.Context().Value("fooID").(string)
	w.Write([]byte(fooID))

}
```