package main

import (
	multi_sort "exercise7.9/multi-sort"
	"html/template"
	"log"
	"net/http"
)

// 练习7.9：利用 html/template 包来替换 printTracks 函数，
// 使用 HTML 表格来显示音乐列表。结合上一个练习，来实现通过单机
// 列头来发送 HTTP 请求，进而对表格排序。

var tracks = []*multi_sort.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, multi_sort.Length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, multi_sort.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, multi_sort.Length("4m24s")},
	{"Go", "Moby", "Moby", 1992, multi_sort.Length("3m37s")},
	{"Go", "Moby2", "Moby2", 1992, multi_sort.Length("3m36s")},
}

var m = &multi_sort.MultiSort{tracks, []string{}}

func main() {
	http.HandleFunc("/", indexHandleFunc)
	http.HandleFunc("/click", clickHandleFunc)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, m.T)
}

func clickHandleFunc(w http.ResponseWriter, r *http.Request) {
	sortType := r.URL.Query().Get("t")
	multi_sort.Click(m, sortType)
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, m.T)
}
