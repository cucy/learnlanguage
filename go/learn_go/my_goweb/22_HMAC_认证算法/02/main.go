package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/authenticate", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")

	// 设置生成新的session, 添加到请求头中
	if err != nil {
		c = &http.Cookie{
			Name:  "session",
			Value: "",
			Path:  "/",
		}
		http.SetCookie(w, c)
	}

	if req.Method == http.MethodPost {
		e := req.FormValue("email")
		c.Value = e + "|" + getCode(e) // 对邮箱进行加密得到加密串
		http.SetCookie(w, c)           // 拼凑 邮箱和 加密的code

	}

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<form method="post">
			<input type="email" name="email">
			<input type="submit">
		</form>
		<a href="/authenticate">Validate This `+c.Value+`</a>
	</body>
	</html>
	`)
}

func auth(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session") // 获取session
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if c.Value == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	xs := strings.Split(c.Value, "|") // 分隔session 得到 邮箱和 加密串
	email := xs[0]
	codeRcvd := xs[1]           // 从客户端得到的值
	codeCheck := getCode(email) // 重新计算值

	fmt.Println(codeRcvd)
	fmt.Println(codeCheck)

	// 匹配两次的值是否一致
	if codeRcvd != codeCheck {
		fmt.Println("HMAC codes didn't match!")
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	io.WriteString(w, `
	<!DOCTYPE html>
	<html>
	<body>
		<h1>`+codeRcvd+` - RECEIVED收到的值 </h1>
		<h1 >`+codeCheck+` - RECALCULATED重新计算的值 </h1>
	</body>
	</html>
	`)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey")) // 秘钥串
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))

}
