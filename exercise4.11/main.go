package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"issue_tool/github"
	"log"
	"os"
)

// 4.11 开发一个工具来让用户可以通过命令行创建、读取、
// 更新或者关闭 GitHub 的 issues，当需要额外输入的时候，
// 调用他们喜欢的文本编辑器。
func main() {
	app := &cli.App{
		Name:  "github issue tool",
		Usage: "you can create, get and update repo issues with this tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "token",
				Usage: "github Token",
			},
			&cli.StringFlag{
				Name:  "repo",
				Usage: "github Repo",
			},
			&cli.StringFlag{
				Name:  "owner",
				Usage: "github Owner",
			},
			&cli.StringFlag{
				Name:  "title",
				Usage: "issue title",
			},
			&cli.StringFlag{
				Name:  "body",
				Usage: "issue body",
			},
			&cli.IntFlag{
				Name:  "number",
				Usage: "issue number",
			},
		},
		Action: func(c *cli.Context) error {
			if c.NArg() > 1 {
				return fmt.Errorf("the number of commands error: %d", c.NArg())
			}

			var (
				repo  = c.String("repo")
				owner = c.String("owner")
				token = c.String("token")

				issue = &github.Issue{
					Repo:  repo,
					Owner: owner,
					Token: token,
				}
			)
			switch c.Args().Get(0) {
			case "create":
				title := c.String("title")
				body := c.String("body")
				issue.Title = title
				issue.Body = body
				issueNumber := issue.Create()
				fmt.Printf("create issue success, the number of issue is: %d \n", issueNumber)
			case "list":
				issues, err := issue.GetList()
				if err != nil {
					fmt.Printf("get list error: %v \n", err)
				}
				for _, i := range issues {
					fmt.Printf("%d, %s, %s, %s", i.Number, i.Title, i.Body, i.State)
				}
			case "get":
				number := c.Int("number")
				issue.Number = number
				i, err := issue.GetOne()
				if err != nil {
					fmt.Printf("get list error: %v \n", err)
				}
				fmt.Printf("%d, %s, %s, %s", i.Number, i.Title, i.Body, i.State)
			case "close":
				number := c.Int("number")
				issue.Number = number

				state := c.String("state")
				issue.State = state
				issue.Close()
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
