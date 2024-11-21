#include <cstdint>
#define i8 char
#define i16 int16_t
#define i32 int
#define i64 long
#define i128 long long
#define u8 char
#define u16 int16_t
#define u32 int
#define u64 long
#define u128 long long
#define alloc(T, S) (T*)malloc(sizeof(T)*S)
#define pbegin Nabs::PFunc profiler_wrap = Nabs::PBegin(__FILE__, __FUNCTION__, __LINE__)
#define pend Nabs::PEnd(profiler_wrap)
#define ppbegin Nabs::PClear()
#define str i8*

#include <stdio.h>
#include "nabs/profiler.hpp"
#
i32 test(i32 a, i32 b) {
pbegin;
pend;
return a+b;
}
int main() {
ppbegin;

i32 i = 0;
for (i32 __UVxeQqsQy = 0; __UVxeQqsQy < 5; __UVxeQqsQy++) { ;
i = test(i, 1);
printf("iterate: %d\n", i);
}
}