package handlers

import (
	"html/template"
	"log"
	"net/http"
	ttemplate "text/template"
	"fmt"
	"os"
)

// 渲染模板
func RenderTemplate(w http.ResponseWriter, templateFile string, templateDate interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)

	}
	t.Execute(w, templateDate)
}

// 不安全的渲染模板
func RenderUnsafeTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := ttemplate.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error encountered while parsing the template: ", err)
	}
	w.Header().Set("X-XSS-Protection", "0")
	t.Execute(w, templateData)
}

// uuid
func GenerateUUID() string {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Println("Encountered the following error when attempting to generate an UUID: ", err)
		return ""
	}
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}
