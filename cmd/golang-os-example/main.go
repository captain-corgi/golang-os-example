package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// Check current OS
	operatingSystem := runtime.GOOS
	fmt.Printf("OS: %s\n", operatingSystem)

	// Check current arch
	architecture := runtime.GOARCH
	fmt.Printf("Arch: %s\n", architecture)

	// List files in current directory as a tree
	if operatingSystem == "windows" {
		listFileTreeWindows()
	} else {
		listFileTreeUnix()
	}

	// Read file
	file, err := os.ReadFile("README.md") // NOTE: Requires go version >= 1.6
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}

// List files in current directory as a tree (Unix)
func listFileTreeUnix() {
	cmd := exec.Command("ls", "-l", "-R", ".")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// List files in current directory as a tree (Windows)
func listFileTreeWindows() {
	cmd := exec.Command("tree", "/f", "/a")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
