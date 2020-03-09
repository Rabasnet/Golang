package main

import (
	"encoding/json"

	"fmt"

	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

//Accont is struct
type Account struct {
	ID          string `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber int64  `json:phonenumber`
	UserName    string `json:"username"`
	PassWord    string `json:"password"`
}

var acc = make(map[string]Account)

func main() {

	inithandler()
}
func inithandler() {
	r := mux.NewRouter()
	r.HandleFunc("/account", createAccountHandler).Methods(http.MethodPost)
	r.HandleFunc("/account", getAccountHandler).Methods(http.MethodGet)
	r.HandleFunc("/account", updateAccountHandler).Methods(http.MethodPut)
	r.HandleFunc("/account", deleteAccountHandler).Methods(http.MethodDelete)

	http.ListenAndServe(":8881", r)
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	var a Account
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	//add new id to the ccount automatically// ="" means nil
	if a.ID == "" {
		a.ID = uuid.New().String()
	}

	// store account details in map
	// map key is id and value is account
	acc[a.ID] = a
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "account created successfully")
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	//read all the acccount parameters

	if id, ok := r.URL.Query()["id"]; ok {
		if err := json.NewEncoder(w).Encode(acc[id[0]]); err != nil {
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		if err := json.NewEncoder(w).Encode(acc); err != nil {
			w.Write([]byte(err.Error()))
			return

		}

	}
	fmt.Fprintln(w, "account read successfully")
}

func updateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var a Account
	if id, ok := r.URL.Query()["id"]; ok {
		//decoding the payload from the request
		if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			fmt.Fprintln(w, "account updated successfully")
			return

		}
		//update the id in the map
		a.ID = id[0]
		//apply the update
		acc[id[0]] = a

		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func deleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.URL.Query()["id"]; ok {
		delete(acc, id[0])

		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
