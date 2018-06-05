
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

# 中间件

## 函数返回函数

```go
func main() {
 numGenerator := generator()
 for i:=0;i<5 ;i++{
 	fmt.Print(numGenerator(),"\t")
 }
}
func generator() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
```

## 自定义中间件

```go
func main() {
	//  HandlerFunc returns a HTTP Handler
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8080", nil)
}

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("执行中间件,在处理request请求之前!")
		//	 pass control back to the handler , 执行handler逻辑
		handler.ServeHTTP(w, r)
		fmt.Println("执行中间件,在处理完response以后")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	//	业务逻辑
	fmt.Println("executing main handler...")
	w.Write([]byte("处理完成,ok!"))
}
```

## 多个中间件

```go

type city struct {
	Name string
	Area uint64
}

// middleware to check content type as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware 当前在检查内容类型中间件中 ")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType) // 不支持的媒体类型 415
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// Middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// Setting cookie to each and every response
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("当前位置处于 Currently in the set server time middlewar  ")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	//	 check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Your resource creation logic goes here. For now it is plain print to console
		// 您的资源创建逻辑在这里。现在是控制台的普通打印。
		log.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
		// Tell everything si fine
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created 创建成功"))

	} else {
		//	 say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
		w.Write([]byte("405 - Method Not Allowed 请求方法不被允许"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":8080", nil)
}
//  curl http://localhost:8080/city  -H "Content-Type:application/json" -iX POST -d '{"Name": "asdfsafdsfa", "Area" : 1000}'

/*
2018/06/04 14:42:10 Currently in the check content type middleware 当前在检查内容类型中间件中
2018/06/04 14:42:10 Got asdfsafdsfa city with area of 1000 sq miles!
2018/06/04 14:42:10 当前位置处于 Currently in the set server time middlewar
*/
```

## multiple Middleware With Alice  多个中间件,链式操作

```go
import(
// ...
	"github.com/justinas/alice"

)
// ....
func main() {
	maincLogincHandler := http.HandlerFunc(mainLogic)
	chain := alice.New(filterContentType,setServerTimeCookie).Then(maincLogincHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8080",nil)
}
```


## 日志中间件

```go

import(
"github.com/gorilla/mux"
"github.com/gorilla/handlers"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request!")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}

```

# RPC

## rpc server

```go
import(
// ...
"net/rpc"
)

type Args struct{}
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	timeserver := new(TimeServer)
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error: ", 1)
	}
	http.Serve(l, nil)
}

```

## rpc client

```go
type Args struct {
}

func main() {
	var reply int64
	args := Args{}
	client, err := rpc.DialHTTP("tcp", "localhost"+":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}
	err = client.Call("TimeServer.GiveServerTime", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	log.Printf("%d", reply)
}
```


## Json rpc server

```
// books.json
[
  {
    "id": "1234",
    "name": "In the sunburned country",
    "author": "Bill Bryson"
  },
  {
    "id":"2345",
    "name": "The picture of Dorian Gray",
    "author": "Oscar Wilde"
  }
]
```

```go
import (
jsonparse "encoding/json"

	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/mux"
)

type Args struct{ Id string }

type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:name,omitempty"`
	Author string `"json:author,omitempty"`
}
type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	raw, readerr := ioutil.ReadFile("./books.json")
	if readerr != nil {
		log.Println("err: ", readerr)
		os.Exit(1)
	}
	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}
	// Iterate over JSON data to find the give book
	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1234", r)
}

// curl http://localhost:1234/rpc  -H "Content-Type:application/json" -iX POST -d '{"method": "JSONServer.GiveBookDetail", "params": [{"Id":"1234"}], "id":1}'

/*
{"result":{"Id":"1234","Name":"In the sunburned country","Author":"Bill Bryson"},"error":null,"id":1}
*/
```


# REST API

go get github.com/emicklei/go-restful

## 最基本的例子


```go
import (
	"github.com/emicklei/go-restful"
// ...
)

func main() {
	werservice := new(restful.WebService)
	werservice.Route(werservice.GET("/ping").To(pingTime))
	restful.Add(werservice)

	http.ListenAndServe(":8080", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
// curl -X GET "http://localhost:8080/ping" -i

2018-06-04 16:36:43.8861582 +0800 CST m=+30.944769901
```

## curd sqlite3

```go
package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Book is a placeholder for book
type Book struct {
	id     int
	name   string
	author string
}

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}
	// Create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER, author VARCHAR(64) NULL, name VARCHAR(64) NULL)")
	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("创建表 Successfully created table books!")
	}
	statement.Exec()
	// Create
	statement, err = db.Prepare("INSERT INTO books (name, author, isbn) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Error in INSERT table", err)
	}
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("插入表 Inserted the book into database!")

	// Read
	rows, err := db.Query("SELECT id, name, author FROM books")
	if err != nil {
		log.Println("Error in Read table", err)
	}
	var tempBook Book
	for rows.Next() {
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("查询表 ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}
	// Update
	statement, err = db.Prepare("update books set name=? where id=?")
	if err != nil {
		log.Println("Error in Update table", err)
	}
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("更新表 Successfully updated the book in database!")

	// Delete
	statement, err = db.Prepare("delete from books where id=?")
	if err != nil {
		log.Println("Error in Delete table", err)
	}
	statement.Exec(1)
	log.Println("删除表 Successfully deleted the book in database!")
}

```



## 创建一个轨道的restfull api

详情 `1_railAPI.md`

1. Design a REST API document.
2. Create models for a database.
3. Implement the API logic.


| HTTP verb | Path                                  | Action | Resource |
| --------- | ------------------------------------- | ------ | -------- |
| POST      | /v1/train (details as JSON body)      | Create | Train    |
| POST      | /v1/station (details as JSON body)    | Create | Station  |
| GET       | /v1/train/id                          | Read   | Train    |
| GET       | /v1/station/id                        | Read   | Station  |
| POST      | /v1/schedule (source and destination) | Create | Route    |

```
curl -X POST http://localhost:8000/v1/trains -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"driverName": "Menaka", "operatingStatus": true}'

curl -X GET "http://localhost:8000/v1/trains/1"
curl -X DELETE "http://localhost:8000/v1/trains/1"
```

# Building RESTful APIs with the Gin framework

`起步`

```go
import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	/*
	GET takes a route and a handler function
	Handler takes the gin context object
*/
	r.GET("/pingTime", func(c *gin.Context) {
		//	 JSON serializer is available on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000")
}
```