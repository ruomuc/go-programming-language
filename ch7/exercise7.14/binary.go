package main

import (
	"fmt"
	"strings"
)

// binary 表示二元操作符表达式，比如 x+y
type binary struct {
	op   rune // '+', '-', '*', '/' 中的一个
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupport binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}

	if err := b.x.Check(vars); err != nil {
		return err
	}
	return nil
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}
