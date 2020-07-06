package eval

// Expr is an arithmetic expression
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

// Var identifies a variable, e.g x, y
type Var string

type literal float64

// unary expression, e.g. -x
type unary struct {
	op rune
	x  Expr
}

// binary expression, e.g. x + y
type binary struct {
	op   rune
	x, y Expr
}

// call expression, e.g. sin(x)
type call struct {
	fn   string
	args []Expr
}

type postUnary struct {
	op rune
	x  Expr
}
