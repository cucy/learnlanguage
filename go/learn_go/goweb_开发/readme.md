
# 简单的http server

```go 

// 罗马数字
var Numerals = map[int]string{
	10: "X",
	9:  "IX",
	8:  "VIII",
	7:  "VII",
	6:  "VI",
	5:  "V",
	4:  "IV",
	3:  "III",
	2:  "II",
	1:  "I",
}

func main() {
	// http包中HandleFunc 来处理请求
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if urlPathElements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404 - Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(Numerals[number]))
			}
		} else {
			// For all other requests, tell that Client sent a bad request
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad request"))
		}

	})

	//	 create a server and  run on 8000 port
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

// http://localhost:8000/roman_number/1
```

# 基本的http handle func

```go 
func main() {
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// hello world, the web server 起码有一个处理函数 
// 接收 w http.ResponseWriter, req *http.Request 参数
func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello,world!")
}
```

## 多个处理函数multiple Handlers

```go 
func main() {
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	http.ListenAndServe(":8000", newMux)
}
```


# 自定义多路径

```go 

type CustomServeMux struct {
}

func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return

}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "随机数字为: %v", rand.Int())

}
func main() {
	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
```

## 使用第三方包路由

```go 
import (
   // ...
	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is:%v\n", vars["category"])
	fmt.Fprintf(w, "Id is:%v\n", vars["id"])
}

func main() {
	// Create a new router
	r := mux.NewRouter()
	// Attach an elegant path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
// http://localhost:8000/articles/1211fsdfs/1221
```

## 获取命令行输出

```go 
import (

	"github.com/julienschmidt/httprouter"
)

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/bin/cat", params.ByName("name")))
}

func getCommandOutput(command string, arguments ...string) string {
	//	 args... unpacks arguments array into elements
	cmd := exec.Command(command, arguments...)

	var (
		out, stderr bytes.Buffer
	)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())

	}
	return out.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("/usr/local/bin/go", "version"))
}

func main() {
	router := httprouter.New()

	// Mapping to methods is possible with HttpRouter
	router.GET("/api/v1/go-version", goVersion)
	// Path variable called name used here
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))

}

```

## 文件服务器

```go 
func main() {
	router := httprouter.New()
	// Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir("/Users/zrd/Desktop/go/go_path/src/learngo"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
```

## 获取url参数

```go 
import (
// ...
	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "Got parameter category:%s!", queryParams["category"][0])
}

func main() {
	// create a new router
	r := mux.NewRouter()

	// Attach an elegant path with handler
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// https://github.com/PacktPublishing/Building-RESTful-Web-Services-with-Go/blob/master/Chapter02/queryParameters.go
// http://localhost:8000/articles?id=0&category=jdnkndssadf1
```
