package eval

import (
	"fmt"
	"math"
	"strings"
)

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}

type call struct {
	fn   string // 'pow', 'sin', 'sqrt' 中的一个
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupport call operator: %q", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, but need %d", c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	var args []string
	for _, a := range c.args {
		args = append(args, a.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(args, ", "))
}

func (c call) Vars() []Var {
	var vars []Var
	for _, a := range c.args {
		vars = append(vars, a.Vars()...)
	}
	return vars
}
