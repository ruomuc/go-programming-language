package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"xkcd/xkcd"
)

// 4.12 流行的Web漫画xkcd有一个JSON接口。例如，调用
// https://xkcd.com/571/info.0.json输出漫画571的详情描述，
// 这个是很多人最喜欢的之一。下载每一个URL并且构建一个离线索引。
// 编写一个工具xkcd来使用这个索引，可以通过命令行指定的搜索条件来查找
// 并输出符合条件的每个漫画的URL和剧本。
func main() {
	app := &cli.App{
		Name:      "xkcd",
		Usage:     "xkcd tool ",
		ArgsUsage: "\n fetch: fetch comics \n search: search comics with keyword",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "min",
				Usage: "fetch comics start number",
			},
			&cli.IntFlag{
				Name:  "max",
				Usage: "fetch comics end number",
			},
			&cli.StringFlag{
				Name:  "k",
				Usage: "the keyword to search comics",
			},
		},
		Action: func(c *cli.Context) (err error) {
			if c.NArg() > 1 {
				return fmt.Errorf("the number of commands error: %d", c.NArg())
			}
			switch c.Args().Get(0) {
			case "fetch":
				min := c.Int("min")
				max := c.Int("max")
				if min < xkcd.Minimum {
					min = xkcd.Minimum
				}
				if max <= min || max > xkcd.Maximum {
					max = xkcd.Maximum
				}
				fmt.Println(min, max)
				xkcd.Fetch(min, max)
			case "search":
				var in *xkcd.Index
				in, err = xkcd.NewIndex()
				if err != nil {
					break
				}
				for _, c := range in.Comics {
					fmt.Printf("title: %s, imgUrl: %s\n", c.Title, c.Img)
				}
			}
			return
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
