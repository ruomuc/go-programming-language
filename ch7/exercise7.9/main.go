package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// 练习7.9：利用 html/template 包来替换 printTracks 函数，
// 使用 HTML 表格来显示音乐列表。结合上一个练习，来实现通过单机
// 列头来发送 HTTP 请求，进而对表格排序。

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go", "Moby2", "Moby2", 1992, length("3m36s")},
}

func main() {

}

func click(m *multiSort, sortField string) {
	m.sortOrder = append(m.sortOrder, sortField)
	sort.Sort(m)
}

type multiSort struct {
	t         []*Track
	sortOrder []string
}

func (m multiSort) Len() int { return len(m.t) }

func (m multiSort) Less(i, j int) bool {
	for _, st := range m.sortOrder {
		switch st {
		case "title":
			if m.t[i].Title != m.t[j].Title {
				return m.t[i].Title < m.t[j].Title
			}
		case "artist":
			if m.t[i].Artist != m.t[j].Artist {
				return m.t[i].Artist < m.t[j].Artist
			}
		case "album":
			if m.t[i].Album != m.t[j].Album {
				return m.t[i].Album < m.t[j].Album
			}
		case "year":
			if m.t[i].Year != m.t[j].Year {
				return m.t[i].Year < m.t[j].Year
			}
		case "length":
			if m.t[i].Length != m.t[j].Length {
				return m.t[i].Length < m.t[j].Length
			}
		default:
			panic(fmt.Sprintf("invalid case value %s", st))
		}
	}
	return false
}

func (m multiSort) Swap(i, j int) { m.t[i], m.t[j] = m.t[j], m.t[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
	fmt.Println()
}
