package main

import (
	"github.com/gogf/gf/g/os/glog"
	"time"
)

func main() {
	SetLog()
}

/*

F_ASYNC      = 1 << iota // 开启日志异步输出
F_FILE_LONG              // 打印调用行号信息，完整绝对路径，例如：/a/b/c/d.go:23
F_FILE_SHORT             // 打印调用行号信息，仅打印文件名，例如：d.go:23，覆盖 F_FILE_LONG.
F_TIME_DATE              // 打印当前日期，如：2009-01-23
F_TIME_TIME              // 打印当前时间，如：01:23:23
F_TIME_MILLI             // 打印当前时间+毫秒，如：01:23:23.675
F_TIME_STD = F_TIME_DATE | F_TIME_MILLI // (默认)打印当前日期+时间+毫秒，如：2009-01-23 01:23:23.675

*/

func SetLog() {
	glog.SetStdoutPrint(false)
	glog.SetFlags(glog.F_TIME_TIME | glog.F_FILE_SHORT)
	path := "./logs"
	glog.SetPath(path)

	glog.Cat("user").Print("测试Print函数 %s", time.Now())
	glog.Cat("user").File("register-{Ymd}.log").Print("测试重命名文件")

	for i := 0; i < 10; i++ {
		glog.Cat("zrd").File("zrd.log").Infof("测试Infof函数 %d", i)

		// 测试json格式
		type User struct {
			Uid  int    `json:"uid"`
			Name string `json:"name"`
		}
		glog.Cat("zrd").File("zrd-json.log").Info(User{Uid: 1, Name: "测试Jsonlog"})

	}

}
