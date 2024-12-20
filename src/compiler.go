package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")

func read_c_integer(value string) int {
	if len(value) == 0 {
		return 0
	}

	if value == "0" {
		return 0
	}

	if value[0] == '0' && len(value) > 1 && value[1] != 'x' && value[1] != 'b' {
		result, err := strconv.ParseInt(value[1:], 8, 64)
		if err != nil {
			return 0
		}
		return int(result)
	}

	if len(value) > 2 && value[:2] == "0x" {
		result, err := strconv.ParseInt(value[2:], 16, 64)
		if err != nil {
			return 0
		}
		return int(result)
	}

	if len(value) > 2 && value[:2] == "0b" {
		result, err := strconv.ParseInt(value[2:], 2, 64)
		if err != nil {
			return 0
		}
		return int(result)
	}

	if len(value) == 3 && value[0] == '\'' && value[2] == '\'' {
		return int(value[1])
	}

	if len(value) > 2 && value[0] == '\'' && value[len(value)-1] == '\'' {
		r, size := utf8.DecodeRuneInString(value[1 : len(value)-1])
		if size > 0 {
			return int(r)
		}
	}
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}

	return result
}

func read_c_integer_expression(value string) int {
	sp := strings.SplitN(value, " ", 3)
	if len(sp) == 1 {
		return read_c_integer(sp[0])
	}
	if len(sp) == 2 {
		println("Invalid expression:", value)
		os.Exit(1)
		return 1
	}
	v1 := read_c_integer(sp[0])
	v2 := read_c_integer(sp[2])
	if sp[1] == "+" {
		return v1 + v2
	}
	if sp[1] == "-" {
		return v1 - v2
	}
	if sp[1] == "*" {
		return v1 * v2
	}
	if sp[1] == "/" {
		if v2 == 0 {
			return 0
		}
		return v1 / v2
	}
	println("Invalid expression:", value)
	os.Exit(1)
	return 0
}

func compile(jpp string) string {
	jpp = "\n\n" + jpp + "\n"
	// jpp = strings.ReplaceAll(jpp, "    ", "  ")
	// jpp = strings.ReplaceAll(jpp, "  ", "\n")
	ec := false
	cpp := bytes.NewBufferString(
		"#ifndef jppt\n" +
			"#define jppt\n" +
			"#define pi(I) printf(\"%d\", I)\n" +
			"#define px(I) printf(\"%02X\", I)\n" +
			"#define pf printf\n" +

			"#ifdef __cplusplus\n" +
			"#include <cstdint>\n" +
			"#else\n" +
			"#include <stdint.h>\n" +
			"#endif\n" +
			"#define i8 char\n" +
			"#define i16 int16_t\n" +
			"#define i32 int\n" +
			"#define i64 long long\n" +
			"#define i128 __int128_t\n" +
			"#define u0 void\n" +
			"#define u8 unsigned char\n" +
			"#define u16 uint16_t\n" +
			"#define u32 unsigned int\n" +
			"#define u64 unsigned long long\n" +
			"#define u128 __uint128_t\n" +
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
			"#define str i8*\n" + // now i8 is c char
			"#endif")

	wrap := 0
	wrapf := 0
	context := 0 // 0 - code, 1 - inline comment, 2 - string, 3 - char, 4 - replaced, 5 - close with )
	// for last_brace := len(jpp) - 1; last_brace >= 0; last_brace-- {
	// 	if jpp[last_brace] == '}' {
	// 		break
	// 	}
	// }
	ms := false
	mms := false
	macro := false
	can_we_maining := true
	for i := 0; i < len(jpp); i++ {
		c := jpp[i]
		// print(string(c))
		l := 0
		if context == 1 && c == '\n' {
			context = 0
		}
		if ec && (c == '{' || c == '\n') {
			cpp.WriteByte(')')
			ec = false
		}
		if c == '(' && context == 5 {
			wrap++
		}
		if c == ')' && context == 5 {
			wrap--
		}
		if context == 2 && c == '"' && jpp[i-1] != '\\' {
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
		if context == 0 && c == '"' && l == 0 && jpp[i-1] != '\\' {
			context = 2
		}
		l = 0
		if c == '\'' && context == 3 && jpp[i-1] != '\\' {
			context = 0
			l = 1
		}
		if c == '\'' && context == 3 && l == 1 && jpp[i-1] != '\\' {
			context = 0
		}
		if context == 2 || context == 1 || context == 4 || context == 5 || context == 3 { // pre-processer loop big wall for defence -----------------------------------------------------------------------------------------------------------------------------250 col
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

		if macro {
			if wrapf == 0 {
				cpp.WriteString("})\n")
				macro = false
				continue
			}
			if c == '\n' {
				cpp.WriteString("\\\n")
				continue
			}
			if !(c == '{' || c == '}' || c == ';' || c == '\n' || c == '/' || c == '\t' || c == '>' || c == ':') &&
				((i < len(jpp)-1) && jpp[i+1] == '\n' ||
					(i < len(jpp)-2 && jpp[i+1] == '/' && jpp[i+2] == '/')) {
				cpp.WriteByte(c)
				cpp.WriteByte(';')
				continue
			}
			cpp.WriteByte(c)
			continue
		}

		if wrapf == 0 && !mms && c != '\n' && c != '{' && c != '}' && c != '#' && i > 0 && can_we_maining {
			w1 := false
			w2 := false
			w3 := false
			n := true
			for j := i; j < len(jpp); j++ {
				if jpp[j] == '\'' && jpp[j-1] != '\\' && !w2 && !w3 {
					w1 = !w1
				}
				if jpp[j] == '"' && jpp[j-1] != '\\' && !w1 && !w3 {
					w2 = !w2
				}
				if jpp[j] == '/' && jpp[j+1] == '/' && !w1 && !w2 {
					w3 = true
					n = false
					break
				}
				if w1 || w2 || w3 {
					continue
				}
				if jpp[j] == '{' || jpp[j] == '#' {
					n = false
					break
				}
				if jpp[j] == '\n' {
					break
				}
			}
			if w1 || w2 || w3 {
				n = false
			}
			if jpp[i-1] == '\n' && jpp[i-2] == '\n' && jpp[i-3] == '\n' && n && !ms {
				cpp.WriteString("i32 main() {\n")
				ms = true
			} else if n && ms {
				cpp.WriteByte(c)
				ms = false
				mms = true
				continue
			}
		}

		if c == '!' && jpp[i+1] == '%' && jpp[i+2] == ' ' && wrapf == 0 && can_we_maining {
			if !ms {
				cpp.WriteString("i32 main() {\n")
			}
			context = 4
			ms = true
			continue
		}
		if c == '!' && jpp[i+1] == '%' && jpp[i+2] == '\n' && wrapf == 0 && can_we_maining {
			cpp.WriteString("}")
			context = 1
			ms = false
			mms = true
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
			if i >= 6 {
				if c == ' ' &&
					jpp[i-1] == 'h' &&
					jpp[i-2] == 'c' &&
					jpp[i-3] == 't' &&
					jpp[i-4] == 'i' &&
					jpp[i-5] == 'w' &&
					jpp[i-6] == 's' {
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
				if strings.HasPrefix(sp, "case(") {
					sP := strings.SplitN(sp[5:][:len(sp)-7], ";", 2)
					v1 := read_c_integer_expression(sP[0])
					v2 := read_c_integer_expression(sP[1])
					println(v1, v2, sP[0], sP[1])
					for j := v1; j <= v2; j++ {
						cpp.WriteString(fmt.Sprintf("case 0x%X: ", j))
					}
					cpp.WriteByte('\n')
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
				if sp == "main" && can_we_maining {
					can_we_maining = false
					cpp.WriteString("int main()")
				}
				if sp == "mainA" && can_we_maining {
					can_we_maining = false
					cpp.WriteString("int main(int argc, char** argv)")
				}
				continue
			}
			if jpp[i-1] == '\n' && c == '#' && jpp[i+1] == '!' {
				s := strings.SplitN(jpp[i+2:], "\n", 2)[0]
				olds := s
				s = strings.ReplaceAll(s, " ", "")
				if strings.HasPrefix(s, "mod:") {
					sp := strings.Split(s[4:], ",")
					for i := 0; i < len(sp); i++ {
						cpp.WriteString("#include \"" + sp[i] + "\"\n")
					}
				}
				if strings.HasPrefix(s, "macro") {
					s := olds
					sp := strings.Split(s[4:], " ")
					if (len(sp) < 4 && len(sp) != 3) || sp[len(sp)-1] != "{" {
						println("correct usage: #!macro MACROSNAME ARG1 ARG2 ... { OR\n#!macro MACROSNAME {")
						os.Exit(1)
					}
					cpp.WriteString("#define " + sp[1])
					if len(sp) != 3 {
						cpp.WriteString("(")
						lsp := len(sp) - 1
						for j := 2; j < lsp; j++ {
							cpp.WriteString(sp[j])
							if j < lsp-1 {
								cpp.WriteString(", ")
							}
						}
						cpp.WriteString(")")
					}
					cpp.WriteString(" (")
					for true {
						if jpp[i] == '{' {
							i--
							break
						}
						i++
					}
					macro = true
					continue
				}
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
			if !(c == '{' || c == '}' || c == ';' || c == '\n' || c == '/' || c == '\t' || c == '>' || c == ':') &&
				((i < len(jpp)-1) && jpp[i+1] == '\n' ||
					(i < len(jpp)-2 && jpp[i+1] == '/' && jpp[i+2] == '/')) {
				cpp.WriteByte(c)
				cpp.WriteByte(';')
				continue
			}
		}
		cpp.WriteByte(c)
	}
	if mms {
		cpp.WriteString("ret 0; \n}")
	}
	return cpp.String()
}
