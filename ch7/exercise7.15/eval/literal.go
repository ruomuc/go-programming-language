package eval

import "fmt"

type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (l literal) Vars() []Var {
	return []Var{Var(l.String())}
}
