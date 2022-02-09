package main

import (
	"exercise2.2/tempconv"
	"fmt"
	"os"
	"strconv"
)

// 练习2.2 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，
// 如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的
// 单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Feet(t)
		p := tempconv.Pound(t)

		fmt.Printf("%s = %s , %s = %s \n", f, tempconv.FToM(f), p, tempconv.PToK(p))
	}
}
