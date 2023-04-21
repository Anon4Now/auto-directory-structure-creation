package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"
	"os"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
	SubDirs SubDir1 `json:"subDir1"`
}

type SubDir1 struct {
	Name string `json:"name"`
	SubDir2 SubDir2 `json:"subDir2"`
}

type SubDir2 struct {
	Name string `json:"name"`
	SubDir3 SubDir3 `json:"subDir3"`
}

type SubDir3 struct {
	Name string `json:"name"`
	SubDir4 SubDir4 `json:"subDir4"`
}

type SubDir4 struct {
	Name string `json:"name"`
	SubDir5 SubDir5 `json:"subDir5"`
}

type SubDir5 struct {
	Name string `json:"name"`
}


func readFile() {
	data, err := os.ReadFile("project.json")

	if err != nil {
		log.Fatal(err)
	}

	var projects Projects

	json.Unmarshal(data, &projects)
	

	// for i := 0; i < len(projects.Projects); i++ {

		
	// }


	// type_of := fmt.Sprintf("%T", data)
	// fmt.Println(type_of)

}

func main() {
	readFile()
}