package main

import (
	"exercise2.1/tempconv"
	"fmt"
)

// 练习2.1  向tempconv包添加类型、常量和函数用来处理Kelvin
// 绝对温度的转换，Kelvin 绝对零度是−273.15°C，Kelvin绝对
// 温度1K和摄氏度1°C的单位间隔是一样的。

const (
	f tempconv.Fahrenheit = 100
	k tempconv.Kelvin     = 100
	c tempconv.Celsius    = 100
)

func main() {
	// KToC
	fmt.Println(tempconv.KToC(k))

	// CToK
	fmt.Println(tempconv.CToK(c))

	// FToK
	fmt.Println(tempconv.FToK(f))

	// KToF
	fmt.Println(tempconv.KToF(k))

}
