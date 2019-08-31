package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Add(var1 int, var2 int) int {
	return var1 + var2
}

func Subtract(var1 int, var2 int) int {
	return var1 - var2
}

func Factorial(num int) int {
	if num > 1 {
		return num * Factorial(num-1)
	} else {
		return 1 // 1! == 1
	}
}

func RootEndpoint (response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))

}

func main() {

	fmt.Println("Starting app...")

	fmt.Println("Add 1+2 : ", Add(1,2))
	fmt.Println("Subtract 10-5 : ", Subtract(10,5))

	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))

}

