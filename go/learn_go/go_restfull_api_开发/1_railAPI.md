
# 文件结构

```
tree  .
.
├── dbutils
│   ├── init-tables.go
│   └── models.go
└── railAPI
    ├── main.go
    └── railapi.db

2 directories, 4 files

```

## models
```go
cat dbutils/models.go
package dbutils

const (
	//  地铁  司机 操作状态
	train = `
CREATE TABLE IF NOT EXISTS train (
ID INTEGER PRIMARY KEY AUTOINCREMENT,
DRIVER_NAME VARCHAR(64) NULL,
OPERATING_STATUS BOOLEAN
)
`
	// 站 站名 起始时间 终止时间
	station = `
CREATE TABLE IF NOT EXISTS station (
ID INTEGER PRIMARY KEY AUTOINCREMENT,
NAME VARCHAR(64) NULL,
OPENING_TIME TIME NULL,
CLOSING_TIME TIME NULL
)
`
	// 调度  地铁id,站id,到达时间,
	schedule = `
CREATE TABLE IF NOT EXISTS schedule (
ID INTEGER PRIMARY KEY AUTOINCREMENT,
TRAIN_ID INT,
STATION_ID INT,
ARRIVAL_TIME TIME,
FOREIGN KEY (TRAIN_ID) REFERENCES train(ID),
FOREIGN KEY (STATION_ID) REFERENCES station(ID)
)
`
)
```

```go
cat dbutils/init-tables.go
package dbutils

import (
	"database/sql"
	"log"
)

func Initialize(dbDriver *sql.DB) {
	statement, driverError := dbDriver.Prepare(train)
	if driverError != nil {
		log.Println(driverError)
	}
	// Create train table
	_, statementError := statement.Exec()
	if statementError != nil {
		log.Println("Table already exists!", statementError)
	}
	statement, statementError = dbDriver.Prepare(station)
	if statementError != nil {
		log.Println("Table already exists!", statementError)
	}
	statement.Exec()
	statement, statementError = dbDriver.Prepare(schedule)
	if statementError != nil {
		log.Println("Table already exists!", statementError)
	}
	statement.Exec()
	log.Println("所有表创建成功 All tables created/initialized successfully!")
}

```

## main

```go
cat railAPI/main.go
package main

import (
	"database/sql"
	"encoding/json"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
	"web/dbutils"
)

// DB Driver visible to whole program 全局变量
var DB *sql.DB

// TrainResource is the model for holding rail information
// 列车资源是铁路信息的保持模型
type TrainResource struct {
	Id              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations 站点信息
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time // 到达时间
}

// Register adds paths and routes to container 注册路径到容器中
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON) // you can specify this per route as well 也可以指定每个路由

	ws.Route(ws.GET("/{train-id}/").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)

}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.Id, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	// Error handling is obvious here. So omitting...错误在这很显而易见, 省略
	statement, _ := DB.Prepare("insert into train (DRIVER_NAME,OPERATING_STATUS) values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.Id = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError,
			err.Error())
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError,
			err.Error())
	}
}

func main() {
	//	连接数据库
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	//	创建表
	dbutils.Initialize(db)

	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on *:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}

```