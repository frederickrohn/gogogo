package main

import (
	"fmt"
	"github.com/frederickrohn/gogogo/user"
	"net/http"
	"encoding/json"
	"sync"
)

type Stats struct {
    TotalUsers    int `json:"total_users"`
    ActiveUsers   int `json:"active_users"`
    ExternalScore int `json:"external_score"`
}

var (
	users = []user.User{
		{ID: 1, Name: "Fred"},
		{ID: 2, Name: "Alice"},
	}

	syncMu = sync.Mutex{}
)

func fetchTotalUsers() int{
	return len(users)
}

func fetchActiveUsers() int{ // just a mock function
	return len(users) / 2
}

func fetchScore() int{ //another mock function
	return 10
}

func statsHandler(w http.ResponseWriter, r* http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	syncMu.Lock()
	total:= fetchTotalUsers()
	active:= fetchActiveUsers()
	syncMu.Unlock()

	score:= fetchScore()
	stats:= Stats{
		TotalUsers: total,
		ActiveUsers: active,
		ExternalScore: score,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func usersHandler(w http.ResponseWriter, r* http.Request){

	switch r.Method{
		case http.MethodGet:

			//lock the mutex and create copy  - faster than doing everything inside the lock
			syncMu.Lock()
			currentUsers:= make([]user.User, len(users))
			copy(currentUsers, users)
			syncMu.Unlock()

			//return the copy of the users slice
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(currentUsers)
		case http.MethodPost:
			// decode json
			var input struct { //look for json key name and assign to Name
				Name string `json:"name"`
			}
			if err:= json.NewDecoder(r.Body).Decode(&input); err != nil{
				http.Error(w, "invalid json", http.StatusBadRequest)
				return
			}
			if input.Name == "" {
				http.Error(w, "name is required", http.StatusBadRequest)
				return
			}

			//create a new user

			syncMu.Lock() //lock when we are changing/reading the users slice
			newID:= len(users) + 1
			newUser:=user.User{
				ID: newID,
				Name: input.Name,
			}
			users = append(users, newUser) //append to users slice
			syncMu.Unlock()

			//return the new user with 201
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) //this returns the 201 status code
			json.NewEncoder(w).Encode(newUser)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	
}



func main() {
	http.HandleFunc("/stats", statsHandler)
	http.HandleFunc("/users", usersHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}