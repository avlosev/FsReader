package main

import (
	"fmt"
	"os"
	"sort"
)

type MyDir struct{
	Name string
	IsDir bool
}

func main() {
	fmt.Println("Goooo!)")
	dir, err := os.Open("D:/Install")
	if err != nil { return }
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil { return }
	//sort.Slice(fileInfos, func(i, j os.FileInfo) bool {
	//	return fileInfos[i].IsDir()<fileInfos[j].IsDir()
	//})
	instDir := []MyDir{}

	for _,file := range fileInfos  {
		//fmt.Println(file.Name(), file.IsDir())
		tmp := MyDir{Name: file.Name(), IsDir: file.IsDir()}
		instDir = append(instDir, tmp)
	}
	sort.Slice(instDir, func(i, j int) bool {
		if instDir[i].IsDir == instDir[j].IsDir {
			return instDir[i].Name<instDir[j].Name
		} else {
			return instDir[i].IsDir
		}
	})
	for _,file := range instDir {
		fmt.Println(file)
	}


	//filepath.Walk("D:/Install", func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(path)
	//	return nil
	//})
}
