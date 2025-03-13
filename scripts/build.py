#!/usr/bin/python3

import os
import subprocess

BUILD_DIR_NAME = "build"

def create_build_directory():
    try:
        if not os.path.exists(BUILD_DIR_NAME):
            os.makedirs(BUILD_DIR_NAME)
    except OSError:
        print("[Error]: Creating directory " + BUILD_DIR_NAME)

def is_go_installed_in_the_system():
    try:
        subprocess.run(["go", "version"], check=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        return True
    except (subprocess.CalledProcessError, FileNotFoundError):
        return False

def build_executables():
    paths = ["cmd/kramer/main.go", "cmd/logjanitor/main.go", "cmd/uncompressor/main.go"]
    for path in paths:
        try:
            os.system("go build -o " + BUILD_DIR_NAME + "/" + path.split("/")[1].split(".")[0] + " " + path)
        except OSError:
            print("[Error]: Building " + path)
    

def main():
    create_build_directory()

    if not is_go_installed_in_the_system():
        print("[Error]: Go is not installed in the system. See: https://go.dev/doc/install")
        return

    print("Building executables...")
    build_executables()
    print("Done!")

if __name__ == "__main__":
    main()
