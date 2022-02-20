package multi_sort

import (
	"fmt"
	"sort"
	"time"
)

func Click(m *MultiSort, sortField string) {
	var exist bool
	for _, so := range m.SortOrder {
		if so == sortField {
			exist = true
		}
	}
	if !exist {
		m.SortOrder = append(m.SortOrder, sortField)
	}
	sort.Sort(m)
}

type MultiSort struct {
	T         []*Track
	SortOrder []string
}

func (m MultiSort) Len() int { return len(m.T) }

func (m MultiSort) Less(i, j int) bool {
	for _, st := range m.SortOrder {
		switch st {
		case "title":
			if m.T[i].Title != m.T[j].Title {
				return m.T[i].Title < m.T[j].Title
			}
		case "artist":
			if m.T[i].Artist != m.T[j].Artist {
				return m.T[i].Artist < m.T[j].Artist
			}
		case "album":
			if m.T[i].Album != m.T[j].Album {
				return m.T[i].Album < m.T[j].Album
			}
		case "year":
			if m.T[i].Year != m.T[j].Year {
				return m.T[i].Year < m.T[j].Year
			}
		case "length":
			if m.T[i].Length != m.T[j].Length {
				return m.T[i].Length < m.T[j].Length
			}
		default:
			panic(fmt.Sprintf("invalid case value %s", st))
		}
	}
	return false
}

func (m MultiSort) Swap(i, j int) { m.T[i], m.T[j] = m.T[j], m.T[i] }

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
