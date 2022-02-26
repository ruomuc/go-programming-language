package eval

import "fmt"

type Expr interface {
	// Eval 返回表达式在 env 上下文下的值
	Eval(env Env) float64

	// Check 方法报告表达式中的错误，并把表达式中的变量假如 Vars 中
	Check(vars map[Var]bool) error

	fmt.Stringer

	Vars() []Var
}

type Env map[Var]float64
