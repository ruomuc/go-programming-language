package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const baseUrl = "https://api.github.com/repos/"

type Issue struct {
	Owner string
	Repo  string
	Token string
	IssueContent
}

type IssueContent struct {
	Number int    `json:"number"`
	State  string `json:"state"`
	IssueBaseContent
}

type IssueBaseContent struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (issue Issue) Create() int {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(issue.IssueBaseContent); err != nil {
		return -1
	}

	u := baseUrl + issue.Owner + "/" + issue.Repo + "/issues"
	r, err := http.NewRequest(http.MethodPost, u, &buf)
	if err != nil {
		fmt.Printf("create issue, http.NewRequest err: %v \n", err)
		return -1
	}
	r.Header.Set("Authorization", "token "+issue.Token)

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("create issue, client.Do error: %v \n", err)
		return -1
	}
	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("create issue, http.Post statusCode not 200, statusCode = %d \n", resp.StatusCode)
		return -1
	}
	defer resp.Body.Close()

	var result IssueContent
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("create issue, resp.body decode err: %v \n", err)
		return -1
	}

	return result.Number
}

func (issue Issue) GetOne() (IssueContent, error) {
	var i IssueContent

	u := baseUrl + issue.Owner + "/" + issue.Repo + "/issues/" + strconv.Itoa(issue.IssueContent.Number)
	resp, err := http.Get(u)
	if err != nil {
		return IssueContent{}, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&i); err != nil {
		return IssueContent{}, err
	}
	return i, nil
}

func (issue Issue) GetList() ([]IssueContent, error) {
	var issues []IssueContent

	u := baseUrl + issue.Owner + "/" + issue.Repo + "/issues"
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return issues, nil
}

func (issue Issue) Close() bool {
	var b bytes.Buffer

	err := json.NewEncoder(&b).Encode(struct {
		State string `json:"state"`
	}{
		State: "close",
	})
	if err != nil {
		fmt.Printf("issue edit, encode err: %v", err)
		return false
	}

	u := baseUrl + issue.Owner + "/" + issue.Repo + "/issues/" + strconv.Itoa(issue.IssueContent.Number)
	r, err := http.NewRequest(http.MethodPatch, u, &b)
	if err != nil {
		fmt.Printf("issue edit, http.NewRequest err: %v", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "token "+issue.Token)

	client := &http.Client{}
	_, err = client.Do(r)
	if err != nil {
		fmt.Printf("issue edit, client.Do err: %v", err)
		return false
	}
	return true
}
