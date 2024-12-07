import os
import sys
if len(sys.argv) != 2:
    print("Usage: make.py <build,test,install,build-all>")
    exit(1)
if sys.argv[1] == "build":
    os.system("go build src/*")
if sys.argv[1] == "build-all":
    if not os.path.exists("build"):
        os.mkdir("build")
    oses = {
        "windows": ["amd64", "386"],
        "darwin": ["amd64"],
        "linux": ["amd64", "arm", "386"]
    }
    extentions = {
        "windows": ".exe",
        "darwin": ".mac",
        "linux": ".elf"
    }
    for OS in oses:
        for ARCH in oses[OS]:
            os.system(f"GOOS={OS} GOARCH={ARCH} go build src/main.go src/compiler.go")
            if OS == "windows":
                os.rename("main.exe", f"build/jpp_{OS}-{ARCH}{extentions[OS]}")
            else:
                os.rename("main", f"build/jpp_{OS}-{ARCH}{extentions[OS]}")
if sys.argv[1] == "test":
    os.system("python3 test.py")

if sys.argv[1] == "install":
    if not os.path.exists("/bin/jpp"):
        os.system("sudo cp build/jpp_linux-amd64.elf /bin/jpp")
        os.system(f"sudo chown -c {os.getuid()} /bin/jpp")
    else:
        os.system("cp build/jpp_linux-amd64.elf /bin/jpp")
    print("pp writed to /bin/jpp")