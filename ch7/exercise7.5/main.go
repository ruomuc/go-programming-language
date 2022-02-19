package main

import (
	"exercise7.5/tempconv"
	"flag"
	"fmt"
)

// 练习7.6：在 tempflag 中支持热力学温度。
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
