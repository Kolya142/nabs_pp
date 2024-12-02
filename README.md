# install:
```
git clone https://github.com/Kolya142/nabs_pp.git
python3 make.py build
python3 make.py install
```

# documentation:
simple macros:
```
i32 - int
u32 - uint

I32 - i32
U32 - u32

u8 - char
str - u8*

nil - NULL
ret - return
alloc(T, S) - alloc S*size(T) bytes with type T
pbegin, pend, ppbegin - nabs profiler, requires #!use:nabsp
```

imports:
```
su - Simple in code Utils (toheap(T, V) - copy V with type T to heap) warning (import std before)
... see compiler.go, tests or examples
```

complex macros:
```
$!main - int main()
$!mainA - int main(int argc, str* argv)
$!lop(T;V;S..E) - iterate from S to E with type T and variable name V
$!lop(C) - repeat code C times
$!loop - for(;;)
#!use:<io,std,ss,ss+,uni,io+,vec,nabsf,nabsp,shash>
```
features:
```
automatic semicolon: "i32 i = 0" -> "i32 i = 0;"
automatic parentheses: "if s == 3 {" -> "if (s == 3) {"
small main: tests/main.jpp -> tests/main.jpp_c.cpp
```