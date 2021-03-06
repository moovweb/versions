package main

import (
	"fmt"
	"github.com/moovweb/versions"
	"os"
)

func usage() {
	fmt.Printf("Usage:\n\tversions <search_path> <name> <version_pattern>\n")
	os.Exit(1)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		usage()
	}

	paths := make([]*versions.FilePath, 0)

	if len(args) == 3 {
		path, err := versions.FindByName(args[1], args[2])

		if err != nil {
			fmt.Println("Error searching for files:\n" + err.Error())
		}

		paths = append(paths, path)
	} else if len(args) == 4 {
		newPaths, err := versions.FindByNameAndVersion(args[1], args[2], args[3])

		if err != nil {
			fmt.Println("Error searching for files:\n" + err.Error())
		}

		paths = newPaths
	} else {
		usage()
	}

	for _, path := range paths {
		fmt.Printf("%v\n", path.String())
	}

}
