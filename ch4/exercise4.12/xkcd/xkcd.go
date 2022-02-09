package xkcd

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const (
	Minimum       = 1
	Maximum       = 2022
	StoreFileName = "xkcd-info.json"
)

type comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

type Index struct {
	Comics []*comic
}

// 下载索引
func Fetch(start, end int) {
	var comics []*comic

	// 获取comic信息
	for i := start; i <= end; i++ {
		comic, err := getComic(i)
		if err != nil {
			fmt.Printf("fetch, get comic err: %v", err)
			break
		}
		comics = append(comics, comic)
	}

	// 离线存储索引
	in := &Index{Comics: comics}
	in.Store()
}

func getComic(num int) (*comic, error) {
	var c comic

	u := "https://xkcd.com/" + strconv.Itoa(num) + "/info.0.json"
	resp, err := http.Get(u)
	if err != nil {
		return nil, errors.Wrap(err, "http get failed")
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&c); err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}
	return &c, nil
}

// 从离线文件读取索引
func NewIndex() (*Index, error) {
	in := &Index{}
	filePtr, err := os.Open(StoreFileName)
	if err != nil {
		return nil, errors.Wrap(err, "open file failed")
	}

	if err = json.NewDecoder(filePtr).Decode(&in.Comics); err != nil {
		return nil, errors.Wrap(err, "json decode failed")
	}

	return in, nil
}

// 离线存储索引
func (in Index) Store() {
	filePtr, err := os.Create(StoreFileName)
	if err != nil {
		fmt.Printf("fetch, creat file err: %v", err)
	}
	defer filePtr.Close()
	if err = json.NewEncoder(filePtr).Encode(in.Comics); err != nil {
		fmt.Printf("store index, json encode err: %v", err)
	}
}

// 从索引中搜索
func (in Index) Search(keyword string) []*comic {
	var comics []*comic
	for _, comic := range in.Comics {
		if match(comic, keyword) {
			comics = append(comics, comic)
		}
	}
	return comics
}

// match 判断comic信息是包含关键字 keyword
func match(c *comic, keyword string) bool {
	return strings.Contains(c.Month, keyword) ||
		strings.Contains(c.Title, keyword) ||
		strings.Contains(c.Alt, keyword) ||
		strings.Contains(c.Day, keyword) ||
		strings.Contains(c.Img, keyword) ||
		strings.Contains(c.Link, keyword) ||
		strings.Contains(c.News, keyword) ||
		strings.Contains(c.SafeTitle, keyword) ||
		strings.Contains(c.Transcript, keyword) ||
		strings.Contains(c.Year, keyword)
}

//{
//"month": "4",
//"num": 571,
//"link": "",
//"year": "2009",
//"news": "",
//"safe_title": "Can't Sleep",
//"transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
//"alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
//"img": "https://imgs.xkcd.com/comics/cant_sleep.png",
//"title": "Can't Sleep",
//"day": "20"
//}
