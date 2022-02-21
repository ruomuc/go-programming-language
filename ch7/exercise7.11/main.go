package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// 练习7.11：增加额外的处理程序，来支持创建、读取、更新和删除数据库条目。比如
// /update?item=socks&price=6这样的请求将更新仓库中物品的价格，，如果商
// 品不存在或者价格无效就返回错误。

var mu sync.RWMutex
var db = database{"shoes": 50, "socks": 5}

func main() {
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/query", db.query)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)

	fmt.Println("server listen on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s \n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s", price)
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	reqQuery := req.URL.Query()
	item, p := reqQuery.Get("item"), reqQuery.Get("price")

	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid param price: %s", p)
		return
	}

	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "item already exist: %s", item)
		return
	}
	// 加写锁
	mu.Lock()
	db[item] = dollars(price)
	// 解写锁
	mu.Unlock()
}

func (db database) query(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s : %s", item, price)
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	reqQuery := req.URL.Query()
	item, p := reqQuery.Get("item"), reqQuery.Get("price")

	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid param price: %s", p)
		return
	}

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
		return
	}
	mu.Lock()
	db[item] = dollars(price)
	mu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	reqQuery := req.URL.Query()
	item := reqQuery.Get("item")

	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %s", item)
		return
	}
	mu.Lock()
	delete(db, item)
	mu.Unlock()
}
