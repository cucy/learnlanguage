package main

import (
	"html/template"
	"log"
	"net/http"

	"fmt"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

type user struct {
	UserName string
	First    string
	Last     string
}

var dbUsers = map[string]user{}      // userID user
var dbSessions = map[string]string{} // sessionId userId

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session") // 从cookie中获取session key
	if err == http.ErrNoCookie {
		sID, _ := uuid.NewV4() // 生成session key
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	// get user if it exists
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		// 获取用户
		u = dbUsers[un]
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un // 保存到数据库中
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session") // 获取
	fmt.Println("用户客户端保存的session:", c)
	if err == http.ErrNoCookie {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]     // 数据库查询
	fmt.Println("un:", un, "Ok:", ok) // un 2@qq.com Ok true
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un] //  结构化值传到模板
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
