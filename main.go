package main

import (

	// "go/token"

	"fmt"
	"interpreter/lexer"
	"interpreter/parser"
	// "interpreter/token"
	// "interpreter/repl"
	// "os"
	// "os/user"
)

func main() {
	// var MonkeySubset string = "let "
	// fmt.Println(MonkeySubset)

	// data, err := os.ReadFile("monkey.txt")

	// if err != nil {
	// 	fmt.Println("Error occured during file reading")
	// 	return
	// }

	// fmt.Println("Content: ")
	// fmt.Println(string(data))

	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Hello %s! This is the Monkey programming language!\n",
	// 	user.Username)
	// fmt.Printf("Feel free to type in commands\n")
	// repl.Start(os.Stdin, os.Stdout)

	// api.Server()

	input := `
	1 + 2 + 3;
	`

	lex := lexer.New(input)

	// for i := 0; i != len(input); i++ {

	// 	tok := lex.NextToken()

	// 	if tok.Type == token.EOF {
	// 		break
	// 	}

	// 	fmt.Println(tok)
	// }

	par := parser.New(lex)
	program := par.ParseProgram()
	fmt.Println(program.String())
	// for i := 0; i != len(program.Statements); i++ {
	// 	fmt.Println(program.TokenLiteral(i))
	// }

}
