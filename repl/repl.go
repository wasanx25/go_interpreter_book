package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/wasanx25/gopter/lexer"
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
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
