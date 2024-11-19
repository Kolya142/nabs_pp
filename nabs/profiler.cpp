#include "profiler.hpp"

namespace Nabs {
    long long timeInMilliseconds(void)  { // https://stackoverflow.com/questions/10192903/time-in-milliseconds-in-c
        struct timeval tv;

        gettimeofday(&tv,nullptr);
        return (((long long)tv.tv_sec)*1000)+(tv.tv_usec/1000);
    }
    PFunc PBegin(const char *fln, const char *fnn, uint32_t lne)
    {
        PFunc pf = PFunc();
        pf.filename = strdup(fln);
        pf.fnname = strdup(fnn);
        pf.line = lne;
        pf.time = timeInMilliseconds();
        FILE* file = fopen("profiler.txt", "a");
        fprintf(file, "function %s at %s:%d started\n", pf.fnname, pf.filename, (int)pf.line);
        fclose(file);
        return pf;
    }
    void PClear()
    {
        FILE* file = fopen("profiler.txt", "w");
        fprintf(file, "Program started\n");
        fclose(file);
    }
    void PEnd(PFunc pf)
    {
        long long tt = timeInMilliseconds();
        int t = tt-pf.time;
        FILE* file = fopen("profiler.txt", "a");
        fprintf(file, "function %s at %s:%d ended, time: %ds, %dms\n", pf.fnname, pf.filename, (int)pf.line, t/1000, t%1000);
        fclose(file);
        free(pf.filename);
        free(pf.fnname);
    }
}