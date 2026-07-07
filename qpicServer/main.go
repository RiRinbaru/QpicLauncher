package main

import (
	"fmt"
	"log"
	"net/http"
	"qpic-server/api"
)

func main() {
	http.HandleFunc("/api/games", api.GetGamesHandler)

	fmt.Println("起動(8080)")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("起動失敗: ", err)
	}
}