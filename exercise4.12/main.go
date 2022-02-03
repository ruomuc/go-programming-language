package main

import (
	"fmt"
	"log"
	"xkcd/xkcd"
)

// 4.12 流行的Web漫画xkcd有一个JSON接口。例如，调用
// https://xkcd.com/571/info.0.json输出漫画571的详情描述，
// 这个是很多人最喜欢的之一。下载每一个URL并且构建一个离线索引。
// 编写一个工具xkcd来使用这个索引，可以通过命令行指定的搜索条件来查找
// 并输出符合条件的每个漫画的URL和剧本。
func main() {
	in, err := xkcd.NewIndex()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	comics := in.Search("p")
	for _, c := range comics {
		fmt.Printf("title: %s, imgUrl: %s\n", c.Title, c.Img)
	}
}
