#define pi(I) printf("%d", I)
#define px(I) printf("%02X", I)
#define pf printf
#ifdef __cplusplus
#include <cstdint>
#else
#include <stdint.h>
#endif
#define i8 char
#define i16 int16_t
#define i32 int
#define i64 long long
#define i128 __int128_t
#define u0 void
#define u8 unsigned char
#define u16 uint16_t
#define u32 unsigned int
#define u64 unsigned long long
#define u128 __uint128_t
#define I8 i8
#define I16 i16
#define I32 i32
#define I64 i64
#define I128 i128
#define U0 u0
#define U8 u8
#define U16 u16
#define U32 u32
#define U64 u64
#define U128 u128
#define nil NULL
#define ret return
#define alloc(T, S) (T*)malloc(sizeof(T)*S)
#define pbegin Nabs::PFunc profiler_wrap = Nabs::PBegin(__FILE__, __FUNCTION__, __LINE__)
#define pend Nabs::PEnd(profiler_wrap)
#define ppbegin Nabs::PClear()
#define str i8*

#include <stdio.h>
#
i32 e(i32 a, i32 b) {
    ret a + b;
}
i32 d(i32 a, i32 b) {
    ret e(a, b)-a+b+b+e(b, a);
}
i32 main() {
pf("Hello, world №%d!\n", e(5032, 1239));
pf("todo: rewrite jpp in jpp\n");
if( e(4, 1) == 5 ){
    pf("checked: 4+1=5\n");
}
pf("d(52, 46) = %d\n", d(52, 46));
ret 0; 
}