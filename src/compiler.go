package main

import (
	"bytes"
	"math/rand"
	"strings"
    "strconv"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")

func compile(jpp string) string {
	jpp = "\n\n" + jpp
	jpp = strings.ReplaceAll(jpp, "    ", "  ")
	jpp = strings.ReplaceAll(jpp, "  ", "\n")
	ec := false
	cpp := bytes.NewBufferString(
		"#ifdef __cplusplus\n" +
		"#include <cstdint>\n" +
		"#else\n" +
		"#include <stdint.h>\n" +
		"#endif\n" +
		"#define i8 int8_t\n" +
		"#define i16 int16_t\n" +
		"#define i32 int\n" +
		"#define i64 long\n" +
		"#define i128 long long\n" +
		"#define u0 void\n" +
		"#define u8 char\n" +
		"#define u16 uint16_t\n" +
		"#define u32 uint\n" +
		"#define u64 ulong\n" +
		"#define u128 ulong long\n" +
		"#define I8 i8\n" +
		"#define I16 i16\n" +
		"#define I32 i32\n" +
		"#define I64 i64\n" +
		"#define I128 i128\n" +
		"#define U0 u0\n" +
		"#define U8 u8\n" +
		"#define U16 u16\n" +
		"#define U32 u32\n" +
		"#define U64 u64\n" +
		"#define U128 u128\n" +
		"#define nil NULL\n" +
		"#define ret return\n" +
		"#define alloc(T, S) (T*)malloc(sizeof(T)*S)\n" +
		"#define pbegin Nabs::PFunc profiler_wrap = Nabs::PBegin(__FILE__, __FUNCTION__, __LINE__)\n" +
		"#define pend Nabs::PEnd(profiler_wrap)\n" +
		"#define ppbegin Nabs::PClear()\n" +
		"#define str u8*\n")

	wrap := 0
	wrapf := 0
	context := 0 // 0 - code, 1 - inline comment, 2 - string, 3 - char, 4 - replaced, 5 - close with )
	ms := false
	for i := 0; i < len(jpp); i++ {
		c := jpp[i]
		l := 0
		if context == 1 && c == '\n' {
			context = 0
		}
		if ec && c == '{' {
			cpp.WriteByte(')')
			ec = false
		}
		if c == '(' && context == 5 {
			wrap++
		}
		if c == ')' && context == 5 {
			wrap--
		}
		if context == 2 && c == '"' {
			context = 0
			l = 1
		}
		if context == 4 && c == ' ' {
			context = 0
		}
		if context == 5 && c == ')' && wrap == 0 {
			context = 0
		}
		if c == '/' && jpp[i+1] == '/' {
			context = 1
		}
		if context == 0 && c == '"' && l == 0 {
			context = 2
		}
		l = 0
		if c == '\'' && context == 3 {
			context = 0
			l = 1
		}
		if c == '\'' && context == 3 && l == 1 {
			context = 0
		}
		if context == 2 || context == 1 || context == 4 || context == 5 || context == 3 { // pre-processer loop wall -----------------------------------------------------------------------------------------------------------------------------250 col
			if context == 2 {
				cpp.WriteByte(c)
			}
			continue
		}
		if c == '{' {
			wrapf++
		}
		if c == '}' {
			wrapf--
		}

		if c == '!' && jpp[i+1] == '%' && jpp[i+2] == ' ' && wrapf == 0 {
			if !ms {
				cpp.WriteString("i32 main() {\n")
			}
			context = 4
			ms = true
			continue
		}
		if c == '!' && jpp[i+1] == '%' && jpp[i+2] == '\n' && wrapf == 0 {
			cpp.WriteString("}")
			context = 1
			continue
		}
		if i != 0 {
			if c == '\n' && jpp[i-1] == '\n' {
				continue
			}
			if i >= 2 {
				if c == ' ' && jpp[i-1] == 'f' && jpp[i-2] == 'i' {
					cpp.WriteByte('(')
					ec = true
				}
			}
			if i >= 3 {
				if c == ' ' && jpp[i-1] == 'r' && jpp[i-2] == 'o' && jpp[i-3] == 'f' {
					cpp.WriteByte('(')
					ec = true
				}
			}
			if i >= 6 {
				if c == ' ' && jpp[i-1] == 'e' && jpp[i-2] == 'l' && jpp[i-3] == 'i' && jpp[i-4] == 'h' && jpp[i-5] == 'w' {
					cpp.WriteByte('(')
					ec = true
				}
			}
			if i < len(jpp)-1 && c == '$' && jpp[i+1] == '!' {
				s := strings.SplitN(jpp[i+1:], " ", 2)[0]

				context = 4

				sp := strings.ReplaceAll(strings.Split(s[1:], " ")[0], " ", "")
				if sp == "loop" {
					cpp.WriteString("for (;;)")
				}
				if strings.HasPrefix(sp, "lop(") {
					sP := strings.SplitN(sp[4:], ";", 3)
					if len(sP) == 0 {
						cpp.WriteString("for (;;)")
					} else if len(sP) == 1 {
						vname := "_"
						for p := 0; p < 10; p++ {
							vname += string(letterRunes[rand.Intn(len(letterRunes))])
						}
						loopLimit := strconv.Itoa(l)
						cpp.WriteString("for (i32 " + vname + " = 0; " + vname + " < " + loopLimit + "; " + vname + "++")
						context = 5
					} else {
						vtype := sP[0]
						vname := sP[1]
						Sp := strings.SplitN(sP[2], "..", 2)
						Sp[1] = Sp[1][:len(Sp[1])-1]
						op := "++"
						op1 := " < "
						a1, b1 := strconv.Atoi(Sp[0])
						a2, b2 := strconv.Atoi(Sp[1])
						if b1 == nil || b2 == nil {
							if a1 > a2 {
								op = "--"
								op1 = " >= "
							}
						}
						cpp.WriteString("for (" + vtype + " " + vname + " = " + Sp[0] + "; " + vname + op1 + Sp[1] + "; " + vname + op)
						context = 5
					}
				}
				if sp == "main" {
					cpp.WriteString("int main()")
				}
				if sp == "mainA" {
					cpp.WriteString("int main(int argc, char** argv)")
				}
				continue
			}
			if jpp[i-1] == '\n' && c == '#' && jpp[i+1] == '!' {
				s := strings.SplitN(jpp[i+2:], "\n", 2)[0]
				s = strings.ReplaceAll(s, " ", "")
				if strings.HasPrefix(s, "use:") {
					sp := strings.Split(s[4:], ",")
					for i := 0; i < len(sp); i++ {
						if sp[i] == "io" {
							cpp.WriteString("#include <stdio.h>\n")
						}
						if sp[i] == "su" {
							cpp.WriteString("#include <string.h>\n")
							cpp.WriteString(
								"void *_toheap(const void *value, U32 size) {\n" +
								"    void *a = malloc(size);\n" +
								"    memcpy(a, value, size);\n" +
								"    return a;\n" +
								"}\n" +
								"#define toheap(T, V) ((T*)_toheap((u0*)&(V), sizeof(V)))\n")
						}
						if sp[i] == "winapi" {
							cpp.WriteString("#include <windows.h>\n")
						}
						if sp[i] == "std" {
							cpp.WriteString("#include <stdlib.h>\n")
						}
						if sp[i] == "math" {
							cpp.WriteString("#include <math.h>\n")
						}
						if sp[i] == "ss" {
							cpp.WriteString("#include <cstring.h>\n")
						}
						if sp[i] == "ss+" {
							cpp.WriteString("#include <string.h>\n")
						}
						if sp[i] == "uni" {
							cpp.WriteString("#include <unistd.h>\n")
						}
						if sp[i] == "io+" {
							cpp.WriteString("#include <iostream>\n")
						}
						if sp[i] == "vec" {
							cpp.WriteString("#include <vector>\n")
						}
						if sp[i] == "nabsf" {
							cpp.WriteString("#include \"nabs/filel.hpp\"\n")
						}
						if sp[i] == "nabsp" {
							cpp.WriteString("#include \"nabs/profiler.hpp\"\n")
						}
						if sp[i] == "shash" {
							cpp.WriteString("#include \"nabs/shash.hpp\"\n")
						}
					}
				}
				context = 1
			}
			if !(c == '{' || c == '}' || c == ';' || c == '\n' || c == '/' || c == '\t' || c == '>') &&
				((i < len(jpp)-1) && jpp[i+1] == '\n' ||
					(i < len(jpp)-2 && jpp[i+1] == '/' && jpp[i+2] == '/')) {
				cpp.WriteByte(c)
				cpp.WriteByte(';')
				continue
			}
		}
		cpp.WriteByte(c)
	}
	return cpp.String()
}
