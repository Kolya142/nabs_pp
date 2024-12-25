#!/bin/python3
import json
import random
import sys
import os
from pathlib import Path

argc = len(sys.argv)
argv = sys.argv
exit = sys.exit

def safe_system(cmd: str) -> None:
    code = os.system(cmd)
    if code != 0:
        print(f"Error: {cmd} exists with {code}.")
        sys.exit(1)

if argc == 2:
    if argv[1] == "init":
        if os.path.exists("lapo_conf.json"):
            print("Error: lapo repo already inited")
            exit(1)
        gargs = ''
        while True:
            inp = input("add g++ args(n/y)?")
            if len(inp) == 0:
                continue
            
            rand = 0
            if inp.replace(' ', '').lower() in \
             'idk;?;i don\' know;i dont know;i do not know;ri;what;'\
             'я не знаю;?;не знаю;не знайу;ри;что;'\
             'издыруам;?;исыздыруам;ри;иарбан;'\
             'ich weiß nicht;?;ich weiß nicht;ich weiß nicht;ich weiß nicht;ri;was;'\
             'わかりません;?;わかりません;わかりません;わかりません;り;何;50%'\
             .lower().replace(' ', '').split(';'):
                if random.randint(0, 100) < 50:
                    rand = 1
                else:
                    rand = 2

            if inp[0].lower() in 'nhнl' or rand == 1:
                break
            if inp[0].lower() in 'yjeдs' or rand == 2:
                while True:
                    inp = input("write args: ")
                    if inp != '':
                        gargs = inp
                        break
                break
        with open("lapo_conf.json", 'w') as f:
            data = {"main": "src/main.jpp", "include": "src/mod", "temp_folder": "target", "output_elf": "main.elf"}
            if gargs:
                data["cargs"] = gargs
            json.dump(data, f)
        safe_system("mkdir src")
        safe_system("mkdir src/mod")
        with open("src/mod/jpp_types.h", 'w') as f:
            f.write(
"""
#ifndef jppt
    #define jppt
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
    #define str i8*
#endif""")
        with open("src/mod/add2.jpp", 'w') as f:
            f.write("#!mod:add2.hpp\ni32 add2(i32 v) {\n    ret v+2\n}\n")
        with open("src/mod/add2.hpp", 'w') as f:
            f.write("#include \"jpp_types.h\"\n#pragma once\ni32 add2(i32 v);")
        with open("src/main.jpp", 'w') as f:
            f.write("#!use: io\n#!mod:add2.hpp\n\n\npf(\"hello, lapo and j++\")\npf(\"\\nalso, 5+2 is %d\\n\", add2(5))\n")
        exit(0)


try:
    conf = json.load(open("lapo_conf.json"))
except FileNotFoundError:
    print("Error: Configuration file 'lapo_conf.json' not found.")
    print(f"Use {sys.argv[0]} init to create minimal lapo setup")
    exit(1)
except json.JSONDecodeError:
    print("Error: Invalid JSON in 'lapo_conf.json'.")
    exit(1)

main = conf["main"]
include = conf["include"]
temp_folder = conf["temp_folder"]
output_elf = conf["output_elf"]
target_type = 'elf' if "target" not in conf else conf["target"]

if not os.path.exists(temp_folder):
    safe_system(f"mkdir {temp_folder}")

files = []

jpps = []

for file in os.listdir(include):
    file: str
    if not file.endswith(".jpp"):
        path = Path(include) / file
        path = str(path)
        pathc = Path(temp_folder) / file
        pathc = str(pathc)
        safe_system(f"cp {path} {pathc}")
        files.append(pathc)
        continue
    jpps.append(file)

plat = 0 if sys.platform.startswith("linux") else (1 if sys.platform.startswith("win") else (2 if sys.platform.startswith("darwin") else -1))
jpp_path = 'jpp' if plat in (0, 2) else 'c:/jpp/jpp.exe'

for jpp in jpps:
    path = Path(include) / jpp
    path = str(path)
    pathc = Path(temp_folder) / (jpp+".cpp")
    pathc = str(pathc)
    safe_system(f"{jpp_path} {path} -o {pathc} -c")
    files.append(pathc)


pathc = Path(temp_folder) / "main.cpp"
pathc = str(pathc)
safe_system(f"{jpp_path} {main} -o {pathc}")
files.append(pathc)

safe_system(f"g++ {' '.join(files)} {'' if 'cargs' not in conf else conf['cargs']} -o {output_elf}")
