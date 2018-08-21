package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wasanx25/gopter/evaluator"
	"github.com/wasanx25/gopter/lexer"
	"github.com/wasanx25/gopter/object"
	"github.com/wasanx25/gopter/parser"
)

const PROMPT = ">> "

const ERROR_AREARE = `
.　　∧∧　∧∧　ｱﾚｱﾚ?
　∩ﾟДﾟ,≡,ﾟДﾟ)
　 ｀ヽ　 　 |）
　　　 | _　|～
　　 　U U
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacro(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, ERROR_AREARE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
