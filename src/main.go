package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		println("Usage: " + os.Args[0] + " <JPP> <CPP> or " + os.Args[0] + " <JPP>")
		os.Exit(0)
	}
	codeb, _ := os.ReadFile(os.Args[1])
	code := string(codeb)
	code = compile(code)
	o := os.Args[1] + "_c.cpp"
	if len(os.Args) == 3 {
		o = os.Args[2]
		os.Exit(0)
	}
	println(os.Args[1], "->", o)
	os.WriteFile(o, []byte(code), 0644)
}
