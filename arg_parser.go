package main

import (
	"flag"
	"log"
)

func parseArgs() []string {
	/* 
	Function that will parse CLI args from user
	*/
	var dirName string
	var dirPath string
	var branches string
	var help bool

	flag.StringVar(&dirName, "name", "default", "(optional) Provide a project name")
	flag.StringVar(&dirPath, "path", "./", "(optional) Provide a path for creation of file structure")
	flag.StringVar(&branches, "branch", "all", "(optional) Choose whether to create with just EPT or IPT second-level directory")
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