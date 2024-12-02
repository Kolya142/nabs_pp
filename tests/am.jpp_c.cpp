#ifdef __cplusplus
#include <cstdint>
#else
#include <stdint.h>
#endif
#define i8 int8_t
#define i16 int16_t
#define i32 int
#define i64 long
#define i128 long long
#define u0 void
#define u8 char
#define u16 uint16_t
#define u32 uint
#define u64 ulong
#define u128 ulong long
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
#define str u8*

#if defined(__x86_64__) || defined(_M_X64) || defined(__aarch64__)
#define mint i64
#define muint u64
#else
#define mint i32
#define muint u32
#endif

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
void *_toheap(const void *value, U32 size) {
    void *a = malloc(size);
    memcpy(a, value, size);
    return a;
}
#define toheap(T, V) ((T*)_toheap((u0*)&(V), sizeof(V)))
#
int main(int argc, char** argv) {
printf("binary name: %s, argument count: %d\n", argv[0], argc);
printf("hello, world!");
i32 i = 0;
i32 v;
if( i <= 0 ){
i = 2;
v = 999;
}
i32* alloced = toheap(i32, v);
printf("memory from heap: %d\n", *alloced);
free(alloced);
for (i32 i = 0; i < 5; i++) {
printf("current iter: %d\n", i);
}
for (i32 _iGEqNbGngI = 0; _iGEqNbGngI < 0; _iGEqNbGngI++) {
printf("current in loop\n");
}
printf("halt program\n");
for (;;) {}
ret 0;
}
