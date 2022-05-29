package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//Ziel:
//Dieses Programm hat den Zweck, alle go files im source folder auszuführen.

const sourcePath = "src/"

func main() {
	files, err := os.ReadDir(sourcePath)
	if err != nil {
		panic(err)
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	wd = strings.ReplaceAll(wd, "\\", "/")

	fileNames := []string{"run"}
	for _, file := range files {
		fileNames = append(fileNames, wd+"/"+sourcePath+file.Name())
	}

	cmd := exec.Command("go", fileNames...)
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		panic(err)
	}
	if err = cmd.Start(); err != nil {
		panic(err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

}
