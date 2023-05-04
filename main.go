package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

// *****MOVE SOME OF THIS TO OTHER FILES*****

func parseArgs() []string {
	/* 
	Function that will parse CLI args from user
	*/
	var dirName string
	var dirPath string
	var branches string
	var help bool

	flag.StringVar(&dirName, "d", "default", "(optional) Provide a project name")
	flag.StringVar(&dirPath, "p", "./", "(optional) Provide a path for creation of file structure")
	flag.StringVar(&branches, "b", "all", "(optional) Choose whether to create with just EPT or IPT second-level directory")
	flag.BoolVar(&help, "help", false, "Help")
	flag.Parse()

	if help {
		flag.PrintDefaults()
	}
	
	if branches != "all" {
		if branches != "EPT" && branches != "IPT" {
			log.Fatal("[-] Bad input for -b, please enter either 'EPT' or 'IPT' : -h for help")
		}
	}
	argSlice := []string{dirName, dirPath, branches}
	return argSlice
}

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

type DirectoryStructure struct {
	/*
	Struct that contains attrs and methods related to the file structure creation.
	*/

	// Instantiate Struct
	projectPath string
	args []string

	// Dynamically Updated
	createdPaths []string
}

func (d *DirectoryStructure) getDataAndResetSlice() []string {
	/*
	Method that will return the data from the struct slice attr and set the slice to nil.
	*/
	data := d.createdPaths
	d.createdPaths = nil
	return data
}

func (d *DirectoryStructure) createSecondLevel() {
	/*
	Method that will build the 'second' level of the directory structure.
	*/
	secondLevel := []string{"EPT", "IPT"}

	// checks for the cli passed from the user around how much to build 
	// NEED TO TEST THE -b PARAMS
	if d.args[2] == "all" {
		for _, a := range secondLevel {
			path := fmt.Sprintf("%v/%v", d.projectPath, a)
			d.createdPaths = append(d.createdPaths, path)
			makeDirCall(path, "mkdirall")
		}
	} else {
		var path string

		if d.args[2] == "EPT" {
			path = fmt.Sprintf("%v/%v", d.projectPath, secondLevel[0])
			d.createdPaths = append(d.createdPaths, path)
		} else if d.args[2] == "IPT" {
			path = fmt.Sprintf("%v/%v", d.projectPath, secondLevel[1])
			d.createdPaths = append(d.createdPaths, path)
		}
	
		makeDirCall(path, "mkdirall")
	}

	log.Println("[+] Successfully created second-level of directory.")
}

func (d *DirectoryStructure) createThirdLevel() {
	/*
	Method that will build the 'third' level of the directory structure.
	*/
	thirdLevel := []string{"evidence", "logs", "scans", "scope", "tools"}

	pathData := d.getDataAndResetSlice()
	fmt.Println(pathData)

	for _, p := range pathData{
		for _, a := range thirdLevel {
			path := fmt.Sprintf("%v/%v", p, a)
			d.createdPaths = append(d.createdPaths, path)
			makeDirCall(path, "mkdirall")
		}
	}

	log.Println("[+] Successfully created third-level of directory.")
}

func (d *DirectoryStructure) createFourthLevel() {
	/*
	Method that will build the 'fourth' level of the directory structure.
	*/
	fourthLevel := []string{"credentials", "data", "screenshots"}

	pathData := d.getDataAndResetSlice()

	for _, p := range pathData {
		for _, a := range fourthLevel {
			if strings.Contains(p, "evidence") {
				path := fmt.Sprintf("%v/%v", p, a)
				makeDirCall(path, "mkdirall")
				continue
			}
		}
	}

	log.Println("[+] Successfully created fourth-level of directory.")

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