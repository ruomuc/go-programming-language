package main

import (
	"issue_tool/github"
	"testing"
	"time"
)

// token自己去账号里新建一个，此token临时测试用，用完就删。
const githubToken = "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"

func TestIssue(t *testing.T) {
	issue := &github.Issue{
		Owner: "ruomuc",
		Repo:  "test_demos",
		Token: githubToken,
		IssueContent: github.IssueContent{
			IssueBaseContent: github.IssueBaseContent{
				Title: "test issue api",
				Body:  "test 2022年2月1日16:27:41",
			},
		},
	}

	t.Run("creat an issue", func(t *testing.T) {
		number := issue.Create()
		if number < 0 {
			t.Errorf("create issue failed")
		}
		issue.IssueContent.Number = number
	})

	time.Sleep(time.Second * 1)

	t.Run("find an issue", func(t *testing.T) {
		got, err := issue.GetOne()
		if err != nil {
			t.Error(err)
		}
		if got.Title != issue.Title || got.Body != issue.Body {
			t.Errorf("got title: %s, got body: %s; but want title: %s, want body: %s", got.Title, got.Body, issue.Title, issue.Body)
		}
	})

	t.Run("find issue list", func(t *testing.T) {
		_, err := issue.GetList()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("update issue", func(t *testing.T) {
		ok := issue.Close()
		if !ok {
			t.Error("update issue failed")
		}
	})
}
