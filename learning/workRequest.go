package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Games struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Company string  `json:"company"`
}

var gamesList = []Games{}

func init() {
	GamesJson := `[{"id":101,"name":"Dragonball","price":1400,"company":"Bandai"},{"id":102,"name":"Naruto","price":1200,"company":"Bandai"},{"id":103,"name":"One Piece","price":1500,"company":"Bandai"}]`
	err := json.Unmarshal([]byte(GamesJson), &gamesList)
	if err != nil {
		log.Fatal(err)
	}
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	gamesJson, err := json.Marshal(gamesList)
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(gamesJson)
	case http.MethodPost:
		var newGame Games
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newGame)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newGame.ID = getNextID()
		gamesList = append(gamesList, newGame)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func getNextID() int {
	newID := -1
	for _, game := range gamesList {
		if game.ID > newID {
			newID = game.ID
		}
	}
	return newID + 1
}

func main() {
	http.HandleFunc("/games", gamesHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
