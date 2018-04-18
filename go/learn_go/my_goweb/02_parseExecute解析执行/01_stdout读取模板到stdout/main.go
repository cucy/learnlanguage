package _1_stdout读取模板到stdout

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
	err = tpl.Execute(os.Stdout, nil) // 执行
	if err != nil {
		log.Fatalln(err)
	}
}
