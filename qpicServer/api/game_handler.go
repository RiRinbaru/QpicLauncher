package api

import (
	"encoding/json"
	"net/http"
	"qpic-server/service"
)

func GetGamesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	games, err := service.GetAllGames()
	if err != nil {
		http.Error(w, "データの取得に失敗した", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(games)
}