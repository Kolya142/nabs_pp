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
i... - ...
u32 - uint
u... - ...
u8 - char
str - u8*
alloc(T, S) - (T*)malloc(typeof(T)*S)
pbegin, pend, ppbegin - nabs profiler, requires #!use:nabsp
```
comple macros:
```
$!main - int main()
$!mainA - int main(int argc, str* argv)
$!for(T;V;S..E) - for (T V = S; V < E; v++)
$!forr(count) - for (i32 _ = 0; _ < count; _++) [_ is random variable name]
#!use:<io,std,ss,ss+,uni,io+,vec,nabsf,nabsp,shash>
```
features:
```
automatic commma: "i32 i = 0" -> "i32 i = 0;"
automatic (): "if s == 3 {" -> "if (s == 3) {"
```