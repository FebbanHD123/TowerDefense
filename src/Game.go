package main

import (
	"fmt"
	"os"
)

var GameFilePath = ""

func main() {
	//Vor.: -
	//Eff.: Start des Spiels

	fmt.Println("Starting Game...")

	fmt.Println("Init Game-Dir")
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	GameFilePath = wd + "/.towerdefense"
	if !FileExistsSafe(GameFilePath) {
		err = os.Mkdir(GameFilePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Game-Dir:", GameFilePath)

	LoadAllLevels()

	initWindow()
	for {
		loop()
	}
}
