package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"os"
	"github.com/bailaohe/go-sql/parser"
	"fmt"
)

type TreeShapeListener struct {
	*parser.BaseSQLiteListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.(type) {
	case *parser.Result_columnContext:
		fmt.Println(ctx.GetText())
	}
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewSQLiteLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewSQLiteParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Sql_stmt_list()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
