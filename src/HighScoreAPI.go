package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const apiServer = "http://45.88.109.123:8080/highscore/v1/"

type Player struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	HighScore int    `json:"highScore"`
}

type a interface{}

func UpdateScoreInAPI(name string, score int) {

	values := map[string]a{"name": name, "score": score}
	jsonData, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(apiServer+"update/", "application/json",
		bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error while send update to api:", err)
		return
	}
	fmt.Println(resp)
}

func GetTopListFromAPI() ([]Player, error) {
	topList := make([]Player, 0)
	resp, err := http.Get(apiServer + "get/toplist/")
	if err != nil {
		fmt.Println("Error while getting the toplist from api:", err)
		return nil, err
	}
	json.NewDecoder(resp.Body).Decode(&topList)
	fmt.Println(topList)
	return topList, nil
}
