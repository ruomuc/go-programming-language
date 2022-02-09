package issue

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const BaseUrl = "https://api.github.com/search/issues"

type IssuesResult struct {
	TotalCount int `json:"total_count"`
	Items      []Issue
}

type Issue struct {
	Title     string
	User      User
	MileStone MileStone
}

type User struct {
	Login string
}

type MileStone struct {
	Title       string
	Description string
}

func SearchIssues(params string) (IssuesResult, error) {
	var res IssuesResult

	q := url.QueryEscape(params)
	resp, err := http.Get(BaseUrl + "?q=" + q)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}
