package main

import (
	"exercise4.10/github"
	"fmt"
	"log"
	"os"
	"time"
)

type category string

const (
	LTONEM = "less than one month"
	LTONEY = "less than one year"
	MTONEY = "more than one year"
)

// 4.10 修改issues实例，按照时间来输出结果，比如一个月以内，
// 一年以内或者超过一年。
func main() {
	issues()
}

func issues() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	categoryMap := make(map[category][]github.Issue)
	for _, item := range result.Items {
		item := *item
		y, m, _ := item.CreatedAt.Date()
		curY, curM, _ := time.Now().Date()

		switch {
		case curY == y && curM-m <= 1:
			categoryMap[LTONEM] = append(categoryMap[LTONEM], item)
		case curY-y <= 1 && curM <= m:
			categoryMap[LTONEY] = append(categoryMap[LTONEY], item)
		default:
			categoryMap[MTONEY] = append(categoryMap[MTONEY], item)
		}
	}

	for c, issues := range categoryMap {
		fmt.Printf("category: %s\n", c)
		for _, issue := range issues {
			fmt.Printf("#%-5d %9.9s %.55s %v\n", issue.Number, issue.User.Login, issue.Title, issue.CreatedAt)
		}
	}
}

// run command
// go run main.go repo:golang/go is:open json decoder

//result
//63 issues:
//#48298     dsnet encoding/json: add Decoder.DisallowDuplicateFields
//#42571     dsnet encoding/json: clarify Decoder.InputOffset semantics
//#43716 ggaaooppe encoding/json: increment byte counter when using decode
//#45628 pgundlach encoding/xml: add Decoder.InputPos
//#48950 Alexander encoding/json: calculate correct SyntaxError.Offset in
//#11046     kurin encoding/json: Decoder internally buffers full input
//#34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
//#36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
//#40128  rogpeppe proposal: encoding/json: garbage-free reading of tokens
//#29035    jaswdr proposal: encoding/json: add error var to compare  the
//#43401  opennota encoding/csv: add Reader.InputOffset method
//#5901        rsc encoding/json: allow per-Encoder/per-Decoder registrati
//#32779       rsc encoding/json: memoize strings during decode
//#30301     zelch encoding/xml: option to treat unknown fields as an erro
//#31701    lr1980 encoding/json: second decode after error impossible
//#28923     mvdan encoding/json: speed up the decoding scanner
//#14750 cyberphon encoding/json: parser ignores the case of member names
//#40982   Segflow encoding/json: use different error type for unknown fie
//#48277 Windsooon encoding/json: add an example for InputOffset() functio
//#48646    piersy encoding/json: unclear documentation for how `json.Unma
//#16212 josharian encoding/json: do all reflect work before decoding
//#6647    btracey x/tools/cmd/godoc: display type kind of each named type
//#28143    arp242 proposal: encoding/json: add "readonly" tag
//#41144 alvaroale encoding/json: Unmarshaler breaks DisallowUnknownFields
//#40127  rogpeppe encoding/json: add Encoder.EncodeToken method
//#34564  mdempsky go/internal/gcimporter: single source of truth for deco
//#33854     Qhesz encoding/json: unmarshal option to treat omitted fields
//#26946    deuill encoding/json: clarify what happens when unmarshaling i
//#22752  buyology proposal: encoding/json: add access to the underlying d
//#43513 Alexander encoding/json: add line number to SyntaxError
