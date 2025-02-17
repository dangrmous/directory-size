package main

import (
	"fmt"
	"github.com/dangrmous/directory-size/directory"
	"github.com/dangrmous/directory-size/format"
	"github.com/dangrmous/directory-size/osfilesystem"
	"log"
	"os"
	"strings"
)

func main() {

	isRecursive := false
	isHuman := false
	baseTen := true

	osf := osfilesystem.OSFileSystem{}

	// Check if the user provided arguments
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory1,directory2,...> --recursive --human", os.Args[0])
	}

	for _, arg := range os.Args[1:] {
		switch arg {
		case "--recursive":
			isRecursive = true
			fmt.Println("Recursive mode enabled")
		case "--human":
			isHuman = true
			fmt.Println("Human readable mode enabled")
		case "--1024":
			baseTen = false
			fmt.Println("JEDEC 100B.01 mode enabled")
		}
	}

	// Get the comma-separated string from the first argument
	dirsArg := os.Args[1]

	// Split the argument into a list of directory names
	dirs := strings.Split(dirsArg, ",")

	// Process each directory
	for _, dir := range dirs {
		// Trim any leading/trailing whitespace from the directory name
		dir = strings.TrimSpace(dir)

		// Skip empty strings (in case there are accidental extra commas)
		if dir == "" {
			continue
		}

		// Get the size of the directory
		dirSize, err := directory.GetDirectorySize(osf, dir, isRecursive)
		if err != nil {
			log.Printf("Error getting size for directory '%s': %v", dir, err)
			continue
		}

		// Print the size of the directory
		if isHuman {
			fmt.Printf("The size of '%s' is %s\n", dir, format.ToHumanReadable(dirSize, baseTen))
		} else {
			fmt.Printf("The size of '%s' is %d bytes\n", dir, dirSize)
		}

	}

}
