package main

import (
	"net/http"
	"html/template"
	"log"

	"go_temp1/vender/socialmedia"
	"fmt"
)

// 显示媒体类型
func displaySocialMediaPostHandler(w http.ResponseWriter, r *http.Request) {
	myPost := socialmedia.NewPost("张三go开发工程师", socialmedia.Moods["thrilled"],
		"Golang 主页!", "看看GO网站！", "https://golang.google.cn",
		"/images/gogopher.png", "", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
	renderTemplate(w, "./templates/socialmediapost.html", myPost)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, tempFile string, tempDate interface{}) {
	// 解析模板
	t, err := template.ParseFiles(tempFile)
	if err != nil {
		log.Fatal("解析模板失败:", err)
	}
	// 执行渲染数据
	t.Execute(w, tempDate)
}

func main() {

	http.HandleFunc("/display-social-media-post", displaySocialMediaPostHandler)
	// 处理静态文件
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
