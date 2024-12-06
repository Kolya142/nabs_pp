package main

import (
	"os";
	"os/exec"
)

func main() {
	if len(os.Args) <= 2 {
		println("Usage: " + os.Args[0] + " <JPP>\noptional flags:\n-o <CPP> - compile to cpp\n-c       - compile to obj\n-e       - compile to elf")
		os.Exit(0)
	}
	codeb, _ := os.ReadFile(os.Args[1])
	code := string(codeb)
	code = compile(code)
	o := os.Args[1] + "_c.cpp"
	co := false
	ce := false

	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "-c" {
			co = true
		}
		if os.Args[i] == "-o" {
			o = os.Args[i+1]
		}
		if os.Args[i] == "-e" {
			ce = true
		}
	}
	println(os.Args[1], "->", o)
	os.WriteFile(o, []byte(code), 0644)
	if ce {
		output, err := exec.Command("g++", o, "-o", o+".elf").CombinedOutput()
		if err != nil {
			println("Compilation error:\n", string(output))
			os.Exit(1)
		}
	}
	if co {
		output, err := exec.Command("g++", "-c", o, "-o", o+".o").CombinedOutput()
		if err != nil {
			println("Compilation error:\n", string(output))
			os.Exit(1)
		}
	}
}
