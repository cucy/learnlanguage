# 使用mux进行路由到不同的handle

```go
import (
	"github.com/gorilla/mux"
	...
)

func main() {
	r := mux.NewRouter()
	http.Handle("/",r)
	http.Handle("/2",r)

	http.ListenAndServe(":8080", nil)
}

```

