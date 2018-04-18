package main

import (
	"net/http"

	"github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	c, err := req.Cookie("session") // 获取session
	if err == http.ErrNoCookie {
		// 如果没有session 则会生成新的session, 设置到请求头中
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	// get the user if it exists 如果存在用户
	var u user
	if un, ok := dbSessions[c.Value]; ok { // session id 是否 获取到 user id  , 否则返回空的用户信息struct结构
		u = dbUser[un]

	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session") // 获取session
	if err != nil {
		return false
	}

	un := dbSessions[c.Value] // 通过 session 获取到 user id
	_, ok := dbUser[un]       // 通过user id 得到  user struct
	return ok
}
