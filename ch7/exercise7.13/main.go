package main

import (
	"fmt"
	"log"
)

// 练习7.13：给Expr增加一个String方法用来美化输出语法树。
// 要求生成的语法树重新解析后是完全一致的树。

func main() {
	tests := []struct {
		expr string
		want string
	}{
		{"-1 + -x", "(-1 + -x)"},
		{"-1 - x", "(-1 - x)"},
		{"sqrt(A / pi)", "sqrt((A / pi))"},
		{"pow(x, 3) + pow(y, 3)", "(pow(x, 3) + pow(y, 3))"},
		{"5 / 9 * (F - 32)", "((5 / 9) * (F - 32))"},
	}
	for i, tt := range tests {
		expr, err := Parse(tt.expr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		got := expr.String()

		if got != tt.want {
			log.Fatalf("%d. got %v, expr %v", i, got, tt.want)
		}
	}
}
