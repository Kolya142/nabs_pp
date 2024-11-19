#pragma once
#include <cstdint>
#include <sys/time.h>
#include <cstdio>
#include <stdlib.h>
#include <string.h>

namespace Nabs {
    long long timeInMilliseconds(void);
    struct PFunc {
        char* filename;
        char* fnname;
        uint32_t line;
        long long time;
    };
    PFunc PBegin(const char* fln, const char* fnn, uint32_t lne);

    void PClear();

    void PEnd(PFunc);
}