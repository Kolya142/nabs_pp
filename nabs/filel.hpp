#include <stdio.h>
#include <cstdio>

namespace Nabs {

    class File {
        FILE* fptr;
        public:
        File(FILE* fptr_);
        char* read();
        void write(char* data);
        void close();
    };

    File open(char* name, char* flags);

}