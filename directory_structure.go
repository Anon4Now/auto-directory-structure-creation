package main

import (
	"fmt"
	"log"
	"strings"
)

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
