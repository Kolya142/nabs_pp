import os
import sys
if len(sys.argv) != 2:
    print("Usage: make.py <build,test,install>")
    exit(1)
if sys.argv[1] == "build":
    os.system("go build src/*")
if sys.argv[1] == "test":
    os.system("python3 test.py")

if sys.argv[1] == "install":
    if not os.path.exists("/bin/jpp"):
        os.system("sudo cp compiler /bin/jpp")
        os.system(f"sudo chown -c {os.getuid()} /bin/jpp")
    else:
        os.system("cp compiler /bin/jpp")
    print("pp writed to /bin/jpp")