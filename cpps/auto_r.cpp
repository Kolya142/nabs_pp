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

if( 3 < 6 ){
x = 3;
}
for( i32 i = 0; i < 5; i++ ){
do something...;
}
while( true ){
1;
}