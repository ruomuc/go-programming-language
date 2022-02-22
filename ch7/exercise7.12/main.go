package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// 练习7.12：修改 /list的处理程序，改为输出 HTML 表格，而不是
// 纯文本。可以考虑使用 html/template 包。

var mu sync.RWMutex
var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", db.list)

	fmt.Println("server listen on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, db); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "tmpl execute err: %s", err)
	}
}
