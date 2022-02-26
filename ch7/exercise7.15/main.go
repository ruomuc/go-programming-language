package main

import (
	"exercise7.15/eval"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// 练习7.15：写一个程序从标准输入读取一个表达式，提示用户输入表达式中的变量的值，
// 最后计算表达式的值。请妥善处理各种异常。

func main() {
	fmt.Print("Expr: \n")

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	expr, err := eval.Parse(string(b))
	if err != nil {
		log.Fatal(err)
	}

	e := env(expr)
	fmt.Printf("Expr = %g\n", expr.Eval(e))
}

func env(expr eval.Expr) eval.Env {
	env := make(eval.Env)

	for _, v := range expr.Vars() {
		var scanVal string

		fmt.Printf("%s: \n", v)
		_, err := fmt.Scanf("%s\n", &scanVal)
		if err != nil {
			log.Fatal("err:", err)
		}

		val, err := strconv.ParseFloat(scanVal, 64)
		if err != nil {
			log.Fatal(err)
		}

		env[v] = val
	}
	return env
}
