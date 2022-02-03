package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"poster/omdb"
)

// 4.13 基于JSON开发的Web服务，开放电影数据库让你可以在
// https://omdbapi.com/上通过名字来搜索电影并下载海报
// 图片。开发一个poster工具以通过命令行指定的电影名称来下载海报。

func main() {
	app := &cli.App{
		Name:  "omdb",
		Usage: "download movie poster from omdb",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "t",
				Usage: "search title of movie",
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("t")
			movie, err := omdb.SearchMovie(title)
			if err != nil {
				return err
			}
			movie.DownLoadPoster()
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
