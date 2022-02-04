package main

import (
	"html/template"
	"issuehtml/issue"
	"log"
	"net/http"
)

// 4.14 创建一个Web服务器，可以通过查询Github并缓存信息，
// 然后可以浏览bug列表、里程碑信息以及参与用户的信息。
// test: http://localhost:8888/?key=is:open%20is:issue%20author:ruomuc%20archived:false
func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}

	q := r.FormValue("key")
	result, err := issue.SearchIssues(q)
	if err != nil {
		log.Fatal(err)
	}
	tmpl := template.Must(template.ParseFiles("issue.html"))
	if err := tmpl.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
