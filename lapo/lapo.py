import json
import os
from pathlib import Path


def safe_system(cmd: str) -> None:
    code = os.system(cmd)
    if code != 0:
        print(f"Error: {cmd} exists with {code}.")
        exit(code)

try:
    conf = json.load(open("lapo_conf.json"))
except FileNotFoundError:
    print("Error: Configuration file 'lapo_conf.json' not found.")
    exit(1)
except json.JSONDecodeError:
    print("Error: Invalid JSON in 'lapo_conf.json'.")
    exit(1)

main = conf["main"]
include = conf["include"]
temp_folder = conf["temp_folder"]
output_elf = conf["output_elf"]

if not os.path.exists(temp_folder):
    safe_system(f"mkdir {temp_folder}")

files = []

for file in os.listdir(include):
    file: str
    if not file.endswith(".jpp"):
        path = Path(include) / file
        path = str(path)
        pathc = Path(temp_folder) / file
        pathc = str(pathc)
        safe_system(f"cp {path} {pathc}")
        continue
    path = Path(include) / file
    path = str(path)
    pathc = Path(temp_folder) / (file+".cpp")
    pathc = str(pathc)
    safe_system(f"jpp {path} -o {pathc} -c")
    files.append(pathc)


pathc = Path(temp_folder) / "main.cpp"
pathc = str(pathc)
safe_system(f"jpp {main} -o {pathc}")
files.append(pathc)

safe_system(f"g++ {' '.join(files)} -o {output_elf}")
