#include <cstdint>
#define i8 int8_t
#define i16 int16_t
#define i32 int
#define i64 long
#define i128 long long
#define u8 char
#define u16 uint16_t
#define u32 uint
#define u64 ulong
#define u128 ulong long
#define alloc(T, S) (T*)malloc(sizeof(T)*S)
#define pbegin Nabs::PFunc profiler_wrap = Nabs::PBegin(__FILE__, __FUNCTION__, __LINE__)
#define pend Nabs::PEnd(profiler_wrap)
#define ppbegin Nabs::PClear()
#define str u8*

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#
void proccessing(bool* state) {
for( i32 i = 25; i > 0; i-- ){
if( state[i-1] ){
state[i] ^= 1;
}
}
}
int main() {
bool* state = alloc(bool, 25);
state[0] = true;
str screen = alloc(u8, 80*25);
memset(screen, ' ', 80*25);
 for (;;) {
proccessing(state);
printf("\x1b[2J\x1b[H");
for( i32 j = 0; j < 25; j++ ){
screen[j*80] = state[25-j] ? '#' : '-';
putchar(screen[j*80]);
for( i32 i = 1; i < 80; i++ ){
screen[i+j*80] = screen[i+j*80+1];
putchar(screen[i+j*80+1]);
}
putchar('\n');
}
usleep(4e4);
}
}