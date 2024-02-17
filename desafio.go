package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Customer struct {
	ID        int     `json:"id,omitempty"`
	FirstName string  `json:"firstName,omitempty"`
	LastName  string  `json:"lastName,omitempty"`
	Email     string  `json:"email,omitempty"`
	Address   Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var customers []Customer

func addCustomer(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	customer.ID = len(customers) + 1
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func getCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, customer := range customers {
		if customer.ID == id {
			json.NewEncoder(w).Encode(customer)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

type homeHandler struct{}

func (h *homeHandler) serverHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bem vinde à página da Vovó!"))
}

func main() {

	home := homeHandler{}

	router := mux.NewRouter()

	// seed data
	customers = append(customers, Customer{
		ID:        1,
		FirstName: "Ana",
		LastName:  "Maria",
		Email:     "ana@email.com",
		Address:   Address{"São Paulo", "SP"},
	})
	customers = append(customers, Customer{
		ID:        2,
		FirstName: "Maria",
		LastName:  "Clara",
		Email:     "maria@email.com",
		Address:   Address{"Belo Horizonte", "MG"},
	})

	router.HandleFunc("/", home.serverHTTP)
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customer/{id}", getCustomerByID).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")

	log.Fatal(http.ListenAndServe(":8010", router))
}
