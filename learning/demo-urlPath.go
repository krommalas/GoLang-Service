package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
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
func gamesHandler2(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "games/")
	gameID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	game, listItemIndex := findID(gameID)
	if game == nil {
		http.Error(w, fmt.Sprintf("no game with id %d", gameID), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		gameJson, err := json.Marshal(game)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJson)
	case http.MethodPut:
		var updateGame Games
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updateGame)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// if updateGame.ID != gameID {
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }
		// game = &updateGame
		// gamesList[listItemIndex] = *game
		//refactor
		updateGame.ID = gameID
		gamesList[listItemIndex] = updateGame
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func findID(ID int) (*Games, int) {
	for i, game := range gamesList {
		if game.ID == ID {
			return &gamesList[i], i
		}
	}
	return nil, 0
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
	http.HandleFunc("/games/", gamesHandler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
