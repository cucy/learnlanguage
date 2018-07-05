package handlers

import (
	"html/template"
	"log"
	"net/http"
	ttemplate "text/template"
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
