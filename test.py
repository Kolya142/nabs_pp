import os

jpps = [i for i in os.listdir("tests") if i.endswith(".jpp")]

for jpp in jpps:
    os.system(f"./compiler tests/{jpp}")
    with open(f"tests/{jpp}_c.cpp") as f:
        print(f.read())
    i = input("is this correct?")
    if i.upper() != 'Y' and i.upper() != "SURE":
        print("so, go fix")
        exit(1)