package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type demo struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type alldemos []demo

var demos = alldemos{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to demoually try it out",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createdemo(w http.ResponseWriter, r *http.Request) {
	var newdemo demo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the demo title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newdemo)
	demos = append(demos, newdemo)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newdemo)
}

func getOnedemo(w http.ResponseWriter, r *http.Request) {
	demoID := mux.Vars(r)["id"]

	for _, singledemo := range demos {
		if singledemo.ID == demoID {
			json.NewEncoder(w).Encode(singledemo)
		}
	}
}

func getAlldemos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(demos)
}

func updatedemo(w http.ResponseWriter, r *http.Request) {
	demoID := mux.Vars(r)["id"]
	var updateddemo demo

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the demo title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updateddemo)

	for i, singledemo := range demos {
		if singledemo.ID == demoID {
			singledemo.Title = updateddemo.Title
			singledemo.Description = updateddemo.Description
			demos = append(demos[:i], singledemo)
			json.NewEncoder(w).Encode(singledemo)
		}
	}
}

func deletedemo(w http.ResponseWriter, r *http.Request) {
	demoID := mux.Vars(r)["id"]

	for i, singledemo := range demos {
		if singledemo.ID == demoID {
			demos = append(demos[:i], demos[i+1:]...)
			fmt.Fprintf(w, "The demo with ID %v has been deleted successfully", demoID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/demo", createdemo).Methods("POST")
	router.HandleFunc("/demos", getAlldemos).Methods("GET")
	router.HandleFunc("/demos/{id}", getOnedemo).Methods("GET")
	router.HandleFunc("/demos/{id}", updatedemo).Methods("PATCH")
	router.HandleFunc("/demos/{id}", deletedemo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
