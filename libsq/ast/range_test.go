package ast

import (
	"testing"

	"github.com/neilotoole/sq/libsq/drvr"
	"github.com/stretchr/testify/assert"
)

// []      select all rows (no range)
// [1]     select row[1]
// [10:15] select rows 10 thru 15
// [0:15]  select rows 0 thru 15
// [:15]   same as above (0 thru 15)
// [10:]   select all rows from 10 onwards

func TestRowRange1(t *testing.T) {

	ast := _getAST(t, FixtRowRange1)
	assert.Equal(t, 0, NewInspector(ast).CountNodes(TypeRowRange))

}

func TestRowRange2(t *testing.T) {

	ast := _getAST(t, FixtRowRange2)
	ins := NewInspector(ast)
	assert.Equal(t, 1, ins.CountNodes(TypeRowRange))
	nodes := ins.FindNodes(TypeRowRange)
	assert.Equal(t, 1, len(nodes))
	rr := nodes[0].(*RowRange)
	assert.Equal(t, 2, rr.Offset)
	assert.Equal(t, 1, rr.Limit)
}

func TestRowRange3(t *testing.T) {

	ast := _getAST(t, FixtRowRange3)
	ins := NewInspector(ast)
	rr := ins.FindNodes(TypeRowRange)[0].(*RowRange)
	assert.Equal(t, 1, rr.Offset)
	assert.Equal(t, 2, rr.Limit)
}

func TestRowRange4(t *testing.T) {

	ast := _getAST(t, FixtRowRange4)
	ins := NewInspector(ast)
	rr := ins.FindNodes(TypeRowRange)[0].(*RowRange)
	assert.Equal(t, 0, rr.Offset)
	assert.Equal(t, 3, rr.Limit)
}

func TestRowRange5(t *testing.T) {

	ast := _getAST(t, FixtRowRange5)
	ins := NewInspector(ast)
	rr := ins.FindNodes(TypeRowRange)[0].(*RowRange)
	assert.Equal(t, 0, rr.Offset)
	assert.Equal(t, 3, rr.Limit)
}
func TestRowRange6(t *testing.T) {

	ast := _getAST(t, FixtRowRange6)
	ins := NewInspector(ast)
	rr := ins.FindNodes(TypeRowRange)[0].(*RowRange)
	assert.Equal(t, 2, rr.Offset)
	assert.Equal(t, -1, rr.Limit)
}

func _getAST(t *testing.T, query string) *AST {
	p := getParser(query)
	q := p.Query()
	ast, err := NewBuilder(drvr.NewSourceSet()).Build(q)
	assert.Nil(t, err)
	assert.NotNil(t, ast)
	return ast
}