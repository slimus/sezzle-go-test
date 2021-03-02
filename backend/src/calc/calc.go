package calc

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

var (
	InvalidExpression = errors.New("expression is invalid")
)

type f func(x, y float64) float64

type Calc struct {
	opMap map[token.Token]f
}

func NewCalc() Calc {
	return Calc{
		opMap: map[token.Token]f{
			token.ADD: func(x, y float64) float64 { return x + y },
			token.SUB: func(x, y float64) float64 { return x - y },
			token.MUL: func(x, y float64) float64 { return x * y },
			token.QUO: func(x, y float64) float64 { return x / y },
		},
	}
}


func (c Calc) Calculate(expression string) (float64, error) {
	e, err := parser.ParseExpr(expression)
	if err != nil {
		return 0, fmt.Errorf("parse expression error %w", err)
	}

	stmts, ok := e.(*ast.BinaryExpr)
	if !ok {
		return 0, fmt.Errorf("can't cast expression to BinaryExpr %w", InvalidExpression)
	}

	return c.calculateTree(stmts)
}

// support only simple expressions
func (c Calc) calculateTree(tree *ast.BinaryExpr) (float64, error) {
	xLit, ok := tree.X.(*ast.BasicLit)
	if !ok {
		return 0, fmt.Errorf("can't convert statement to literal")
	}
	x, err := strconv.ParseFloat(xLit.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("can't cast string to float64 %w", err)
	}

	yLit, ok := tree.Y.(*ast.BasicLit)
	if !ok {
		return 0, fmt.Errorf("can't convert statement to literal")
	}
	y, err := strconv.ParseFloat(yLit.Value, 64)
	if err != nil {
		return 0, fmt.Errorf("can't cast string to float64 %w", err)
	}

	return c.opMap[tree.Op](x, y), nil
}
