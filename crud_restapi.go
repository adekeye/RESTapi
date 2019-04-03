package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Site struct
type Site struct {
	Name        string       `json:"name,omitempty"`
	Role        string       `json:"role,omitempty"`
	URI         string       `json:"uri,omitempty"`
	AccessPoint *AccessPoint `json:"AccessPoint,omitempty"`
}

// AccessPoint struct
type AccessPoint struct {
	Label string `json:"Label,omitempty"`
	URL   string `json:"URL,omitempty"`
}

// Initialize Sites variable
var Sites []Site

// Get all Sites
func getSites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Sites)
}

// Get single Site
func getSite(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req) // Gets params
	// Loop through Sites and find one with the Name from the params
	for _, element := range Sites {
		if element.Name == params["Name"] {
			json.NewEncoder(w).Encode(element)
			return
		}
	}
	json.NewEncoder(w).Encode(&Site{})
}

// Add new Site
func createSite(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	var Site Site
	_ = json.NewDecoder(req.Body).Decode(&Site)
	Site.Name = params["Name"]
	Sites = append(Sites, Site)
	json.NewEncoder(w).Encode(Site)
}

// Update Site
func updateSite(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, element := range Sites {
		if element.Name == params["Name"] {
			Sites = append(Sites[:index], Sites[index+1:]...)
			var Site Site
			_ = json.NewDecoder(req.Body).Decode(&Site)
			Site.Name = params["Name"]
			Sites = append(Sites, Site)
			json.NewEncoder(w).Encode(Site)
			return
		}
	}
}

// Delete Site
func deleteSite(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, element := range Sites {
		if element.Name == params["Name"] {
			Sites = append(Sites[:index], Sites[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Sites)
}

// Main function
func main() {
	// Initialize router
	r := mux.NewRouter()
	//Sample Data
	Sites = append(Sites, Site{Name: "Google", Role: "Search", URI: "138.292.112", AccessPoint: &AccessPoint{Label: "Main", URL: "www.google.com"}})
	Sites = append(Sites, Site{Name: "MyUmbc", Role: "Student Services", URI: "121.123.221", AccessPoint: &AccessPoint{Label: "main", URL: "www.myumbc3.edu"}})

	//Sites = append(Sites, Site{Name: "C://", Role: "disk", URI: "180:202:112", AccessPoint: &AccessPoint":[]})

	r.HandleFunc("/Sites", getSites).Methods("GET")
	r.HandleFunc("/Sites/{Name}", getSite).Methods("GET")
	r.HandleFunc("/Sites", createSite).Methods("POST")
	r.HandleFunc("/Sites/{Name}", updateSite).Methods("PUT")
	r.HandleFunc("/Sites/{Name}", deleteSite).Methods("DELETE")

	// Server Init
	log.Fatal(http.ListenAndServe(":8000", r))
}
