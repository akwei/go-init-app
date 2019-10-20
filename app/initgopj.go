package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

/*
create standard golang app project

intput flag
	-prefix			project parent path
	-name			project name

*/

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	subpath := make([]string, 0, 20)
	subpath = append(subpath, "api")
	subpath = append(subpath, "assets")
	subpath = append(subpath, "build")
	subpath = append(subpath, "cmd")
	subpath = append(subpath, "configs")
	subpath = append(subpath, "deployments")
	subpath = append(subpath, "docs")
	subpath = append(subpath, "examples")
	subpath = append(subpath, "githooks")
	subpath = append(subpath, "init")
	subpath = append(subpath, "internal")
	subpath = append(subpath, "pkg")
	subpath = append(subpath, "scripts")
	subpath = append(subpath, "test")
	subpath = append(subpath, "third_party")
	subpath = append(subpath, "tools")
	subpath = append(subpath, "vendor")
	subpath = append(subpath, "web")
	subpath = append(subpath, "website")

	// get flag
	_prefix := flag.String("prefix", "", "prefix must be set. A path that project be created")
	_name := flag.String("name", "", "name must be set. Project name")
	flag.Parse()

	args := os.Args
	fmt.Println(args)

	prefix := *_prefix
	name := *_name

	// -prefix=/Users/akwei/project_go/auto_create/ -name=auto1
	//prefix := "/Users/akwei/project_go/auto_create/"
	//name := "auto1"

	if len(prefix) == 0 {
		panic("prefix must be set")
	}
	if len(name) == 0 {
		panic("name must be set")
	}

	// create parent path
	p := filepath.Join(prefix, name)
	_, e := os.Stat(p)
	if os.IsExist(e) {
		checkErr(e)
	}
	e = os.MkdirAll(p, 0755)
	checkErr(e)

	// create subDir
	for idx, subDir := range subpath {
		subDirPath := filepath.Join(prefix, name, subDir)
		fmt.Printf("%d create dir: %s\n", idx, subDirPath)
		e := os.MkdirAll(subDirPath, 0755)
		checkErr(e)
	}
	fmt.Println("finish create project")
}
