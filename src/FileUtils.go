package main

import "os"

func FileExists(path string) (bool, error) {
	//Vor.: path wird als parameter übergeben
	//Eff.: Es wird zurückgegeben, ob auf dem path ein File/Dir existiert oder nicht
	//Allerdings wird, falls es ein Fehler gibt, dieser ebenfalls zurückgegeben
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func FileExistsSafe(path string) bool {
	//Vor.: path wird als parameter übergeben
	//Eff.: Es wird zurückgegeben, ob auf dem path ein File/Dir existiert oder nicht
	exists, err := FileExists(path)
	if err != nil {
		panic(err)
	}
	return exists
}
