package main

import "fmt"

type ternary struct {
	op1, op2 rune // '?' or ':'
	x, y, z  Expr
}

func (t ternary) Eval(env Env) float64 {
	if t.x.Eval(env) != 0 {
		return t.y.Eval(env)
	}
	return t.z.Eval(env)
}

func (t ternary) Check(vars map[Var]bool) error {
	if t.op1 != '?' || t.op2 != ':' {
		return fmt.Errorf("unexpected ternary op : %q , %q", t.op1, t.op2)
	}

	if err := t.x.Check(vars); err != nil {
		return err
	} else if err := t.y.Check(vars); err != nil {
		return err
	}
	return t.z.Check(vars)
}

func (t ternary) String() string {
	return fmt.Sprintf("%s %c %s %c %s", t.x, t.op1, t.y, t.op2, t.z)
}
