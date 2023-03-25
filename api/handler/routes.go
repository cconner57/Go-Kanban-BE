package routes

import (
	"encoding/json"
	"kanban/api/types"

	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"log"
)

var db *sql.DB

func GetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	rowsRs, err := db.Query("SELECT * FROM boards")

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	defer rowsRs.Close()

	var boards []types.Board = []types.Board{}

	for rowsRs.Next() {
		brd := types.Board{}
		err := rowsRs.Scan(&brd.Id, &brd.Name)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		boards = append(boards, brd)
	}

	if err = rowsRs.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, brd := range boards {
		fmt.Fprintf(w, "%s %s\n", brd.Id, brd.Name)
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(boards)	
}