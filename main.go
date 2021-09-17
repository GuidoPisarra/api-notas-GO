package main

//gorilla mux es un enrutadorDefinir middleware en las rutas, es decir, aplicar funciones
//que se ejecutan antes de cada petición HTTP y que permiten detener la ejecución o loguear determinadas cosas.
// Definición de rutas con verbos HTTP. Lectura de parámetros GET
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//importo modulo para convertir string en int
	"strconv"
	//maneja E/S del servidor
	"io/ioutil"
	//"log"
	//mux es el enrutador
	"github.com/gorilla/mux"
)

//Defino Persona
type Person struct {
	// le digo que puede recibir en formato json y que no tenga en cuenta a los vacios
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	// con el * le digo que apunte al tipo de dato Address
	Address *Address `json:"address,omitempty"`
}
type Address struct {
	State  string `json:"state,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Nro    int    `json:"nro,omitempty"`
}

//mock de prueba
var people []Person

// la w es de write es decir que puede "escribir " una respuesta
// response writer es el que devuelve el arreglo de datos
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mi Primer  API con  GO")

}

func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "application/json")

	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "id invalido")
		return
	}
	for _, perso := range people {
		if perso.ID == personID {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(perso)
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	reqBody, err := ioutil.ReadAll(req.Body)
	var newPerson Person
	if err != nil {
		fmt.Fprintf(w, "Inserte datos validos de una persona")
	}
	json.Unmarshal(reqBody, &newPerson)
	newPerson.ID = len(people) + 1
	people = append(people, newPerson)
	//envio info extra por cada peticion
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprintf(w, "id invalido, no se pudo eliminar")
		return
	}
	for i, p := range people {
		if p.ID == personID {
			people = append(people[:i], people[i+1:]...)
			fmt.Fprintf(w, "la persona ID %v fue eliminada", personID)
		}
	}
	json.NewEncoder(w).Encode(people)
}

func UpdatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	personID, err := strconv.Atoi(vars["id"])
	var updatePerson Person
	if err != nil {
		fmt.Fprintf(w, "id invalido, no se puede modificar")
		return
	}
	reqbody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "ingrese datos validos")
	}
	json.Unmarshal(reqbody, &updatePerson)
	for i, p := range people {
		if p.ID == personID {
			people = append(people[:i], people[i+1:]...)
			updatePerson.ID = personID
			people = append(people, updatePerson)
			fmt.Fprintf(w, "la persona ID %v fue modificada con exito", personID)

		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	//routes
	router := mux.NewRouter()

	//mock de prueba
	people = append(people, Person{ID: 1, FirstName: "Juan", LastName: "Gomez", Address: &Address{State: "Buenos Aires", City: "Azul", Street: "San Martin", Nro: 888}})
	people = append(people, Person{ID: 2, FirstName: "Pepe", LastName: "Grillo", Address: &Address{State: "Buenos Aires", City: "Tres Arroyos", Street: "Colon", Nro: 333}})

	//endpoints
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/create", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	router.HandleFunc("/people/modify/{id}", UpdatePersonEndPoint).Methods("PUT")
	//creo el srvidor
	http.ListenAndServe(":3000", router)
	//log es el modulo que me informa los errores en caso de que los haya
	log.Fatal(http.ListenAndServe(":3000", nil))
}
