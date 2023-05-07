package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)


func makeDirCall(path string, method string) {
	/*
	Function that will call the os.Mkdir/os.MkdirAll methods based on arg.
	*/
	commands := map[string]func(string, fs.FileMode) error {
		"mkdir": os.Mkdir,
		"mkdirall": os.MkdirAll,
	}
	if method != "mkdir" && method != "mkdirall" {
		log.Fatal("[-] Error processing sent os method.")
	}

	err := commands[method](path, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}


func main() {
	args := parseArgs()

	log.Printf("[!] Attempting to create a new project '%v' at location '%v'\n", args[0], args[1])

	projectPath := fmt.Sprintf("%v/%v", args[1], args[0])

	// make the root directory for the project
	makeDirCall(projectPath, "mkdir")

	dir := DirectoryStructure{projectPath: projectPath, args: args}

	dir.createSecondLevel()
	dir.createThirdLevel()
	dir.createFourthLevel()
}