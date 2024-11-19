#include "filel.hpp"
#include <vector>
#include <string>
#include <iostream>
#include <string.h>

namespace Nabs {

    File open(char* name, char* flags) {
        File f = File(fopen(name, flags));
        return f;
    }
    File::File(FILE* fptr_) {
        fptr = fptr_;
    }
    char* File::read() {
        std::string str = std::string();
        char* chunk = (char*)malloc(256);
        while (fgets(chunk, 256, fptr)) {
            str += chunk;
        }
        free(chunk);
        char* result = (char*)(malloc(str.size() + 1));
        strcpy(result, str.c_str());
        return result;
    }
    void File::write(char* data) {
        fprintf(fptr, data);
    }
    void File::close() {
        fclose(fptr);
    }
}