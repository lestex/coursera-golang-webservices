package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type DirStructure struct {
	Name  string
	IsDir bool
	Size  int64
	Items []DirStructure
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"

	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	var parent DirStructure

	parent.Name = path
	parent.IsDir = true
	parent.Items = createDirStructure(path)

	out1 := PrintTree(parent)
	fmt.Fprintln(out, out1)
	return nil
}

func createDirStructure(directory string) []DirStructure {

	var items []DirStructure
	files, _ := ioutil.ReadDir(directory)

	for _, f := range files {

		var child DirStructure
		child.Name = f.Name()

		if f.IsDir() {
			newDirectory := filepath.Join(directory, f.Name())
			child.IsDir = true
			child.Items = createDirStructure(newDirectory)
		} else {
			child.Size = f.Size()
		}

		items = append(items, child)
	}
	return items
}

func StringTree(object DirStructure) (result string) {
	var spaces []bool
	result += stringObjItems(object.Items, spaces)
	return
}

func stringLine(name string, spaces []bool, last bool) (result string) {
	for _, space := range spaces {
		if space {
			result += "	"
		} else {
			result += "│	"
		}
	}

	indicator := "├───"
	if last {
		indicator = "└───"
	}

	result += indicator + name + "\n"
	return
}

func stringObjItems(items []DirStructure, spaces []bool) (result string) {
	for i, f := range items {
		last := (i >= len(items)-1)
		result += stringLine(f.Name, spaces, last)
		if len(f.Items) > 0 {
			spacesChild := append(spaces, last)
			result += stringObjItems(f.Items, spacesChild)
		}
	}
	return
}

func PrintTree(object DirStructure) string {
	return StringTree(object)
}

func PrintFileSize() {

}
