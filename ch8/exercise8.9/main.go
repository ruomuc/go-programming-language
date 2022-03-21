package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。

var (
	wg      sync.WaitGroup
	verbose = flag.Bool("v", false, "show verbose progress messages")
)

type dir struct {
	idx       int
	name      string
	filesSize int64
}

func main() {
	// 确定初始目录
	flag.Parse()
	roots := os.Args[1:]
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 遍历文件树
	dirsChan := make(chan dir)
	for idx, root := range roots {
		wg.Add(1)
		go walkDir(root, idx, dirsChan)
	}

	go func() {
		wg.Wait()
		close(dirsChan)
	}()

	// 输出结果
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case d, ok := <-dirsChan:
			if !ok {
				break loop
			}
			nfiles[d.idx]++
			nbytes[d.idx] += d.filesSize
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}
	printDiskUsage(roots, nfiles, nbytes)
}

func walkDir(dirName string, idx int, dirsChan chan<- dir) {
	defer wg.Done()

	for _, entry := range dirents(dirName) {
		if entry.IsDir() {
			wg.Add(1)

			subDir := filepath.Join(dirName, entry.Name())
			walkDir(subDir, idx, dirsChan)
		} else {
			dirsChan <- dir{idx, dirName, entry.Size()}
		}
	}
}

// sema是一个用于限制目录并发数的计数信号量
var sema = make(chan struct{}, 20)

func dirents(dirName string) []os.FileInfo {
	sema <- struct{}{}        // 获取令牌
	defer func() { <-sema }() // 释放令牌

	entries, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for idx, dirName := range roots {
		fmt.Printf("%s: %d files %.1f MB\n", dirName, nfiles[idx], float64(nbytes[idx]/1e6))
	}
}
