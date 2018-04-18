package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // 解析
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("创建文件失败", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil) // 执行   输出到文件
	if err != nil {
		log.Fatalln(err)
	}
}
