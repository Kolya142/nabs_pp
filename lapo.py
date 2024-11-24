#!/bin/python3
import pathlib
import json
import os

# Defines 
config = None
error = "\x1b[31mERROR\x1b[0m"
warning = "\x1b[33mWARNING\x1b[0m"

# Json parse

if not os.path.exists("config.json"):
    print("failed to found config.json")
    with open("config.json", 'w') as f:
        f.write(
"""{
"language": "c++",
"compiler_args": "",
"includes": {
    "jppd": ["src/"],
    "jppf": [],
    "cppd": [],
    "cppf": []
},
"output_dir": "cpps",
"macros": []
}""")
    print("config.json created")
    exit(1)

with open("config.json") as f:
    config = json.load(f)
    has = {"language": False, "compiler_args": False, "includes": False, "output_dir": False, "macros": False}
    for subobj in config:
        if subobj not in has:
            continue
        if has[subobj]:
            print(error + f": {subobj}")
            exit(1)
        has[subobj] = True
    if not all(has.values()):
        print(error + ": not all properties set in config file")
        exit(1)

# decode json object

lang = config["language"]
cargs = config["compiler_args"]
includes_jppd = config["includes"]["jppd"]
includes_jppf = config["includes"]["jppf"]
includes_cppd = config["includes"]["cppd"]
includes_cppf = config["includes"]["cppf"]
output = config["output_dir"]
macros = config["macros"]
if macros:
    print(warning+": macros not realised")
compiler = "g++" if lang == "c++" else "gcc"

if not os.path.exists(output):
    os.mkdir(output)

# Copy c++ files
if includes_cppd:
    if os.system(f'cp {" ".join(i + "/*" for i in includes_cppd)} {output}') != 0:
        print(error + ": one of submodules gets error")
        exit(1)
if includes_cppf:
    if os.system(f'cp {" ".join(includes_cppf)} {output}') != 0:
        print(error + ": one of submodules gets error")
        exit(1)

# Compile jpp
jpps: list = includes_jppf
for jppd in includes_jppd:
    for d in os.listdir(jppd):
        jpps.append(os.path.join(jppd, d))
print(jpps)
for jpp in jpps:
    o = f"{output}/{pathlib.Path(jpp).name.split('.')[0]}.cpp"
    print(f"compiling: {jpp} to {o}")
    cmd = f"jpp {jpp}"
    print(cmd)
    if os.system(cmd) != 0:
        print(error + ": one of submodules gets error")
        exit(1)
    os.system(f"mv {jpp}_c.cpp {o}")
    if not os.path.exists(o):
        print(error + ": jpp file not compiled")
        exit(1)

# Compile c++

if os.system(f'{compiler} {output}/*.cpp -o {output}.o {cargs}') != 0:
    print(error + ": one of submodules gets error")
    exit(1)

