import os

srcs = ""
for root,dirs,files in os.walk(os.path.join(os.getcwd(), "cmd")):
    for file in files:
        srcs += os.path.join(root, file)+" "
os.system("go run "+srcs)
