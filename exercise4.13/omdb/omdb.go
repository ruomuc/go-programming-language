package omdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

// 这个apikey可能会用不了，需要自己去这个网站上申请。
const BaseURL = "http://www.omdbapi.com/?apikey=b5cdf3a8"

type Movie struct {
	Title  string
	Poster string
}

func (m Movie) DownLoadPoster() {
	poster := m.Poster
	suffix := path.Ext(poster)
	if suffix == "" {
		suffix = ".jpg"
	}

	name := strings.ReplaceAll(m.Title, " ", "")
	name = strings.ReplaceAll(name, ":", "")
	imgName := name + suffix

	resp, err := http.Get(poster)
	if err != nil {
		fmt.Printf("DownLoadPoster, http.Get err: %v", err)
		return
	}
	defer resp.Body.Close()

	out, _ := os.Create(imgName)
	defer out.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	io.Copy(out, bytes.NewReader(body))
	return
}

func SearchMovie(title string) (Movie, error) {
	var movie = Movie{}

	t := url.QueryEscape(title)
	u := BaseURL + "&t=" + t + "&plot=short&r=json"

	resp, err := http.Get(u)
	if err != nil {
		return movie, nil
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return movie, err
	}
	return movie, nil
}
