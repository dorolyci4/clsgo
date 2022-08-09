import os

# go test ./test -v
# -run=none means only benchmark test executed
# -benchtime=10s means time to running benchmark
# go test ./test -v -bench=. -memprofile memprofile.out -cpuprofile profile.out

for root,dirs,files in os.walk(os.path.join(os.getcwd(), "pkg")):
    for dir in dirs:
        print(os.path.join(root, dir))
        os.system("go test "+os.path.join(root, dir)+" -v ")

os.system("go test ./test -v -bench=\"./test\"")