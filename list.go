// package main

// import "fmt"

// var names []string = []string{"Dean", "Myranda", "Will"}

// func main() {

// 	fmt.Println(names[1])
// }

// package main

// import "fmt"

// func main() {
// var name string = "DeanKeene"
// fmt.Printf("My name is %v!\n", name)

// startingPrice := 29.99
// fmt.Printf("Starting price begins at %v\n", startingPrice)
// fmt.Println("After 10% discount, the price is:", startingPrice*0.9)

// temperatureK := 288.15
// temperatureF := (temperatureK-273.15)*1.8 + 32

// fmt.Printf("The Temperature is %v\n", temperatureF)

// var temperatureC float64 = 10
// fmt.Printf("Integer to float: %.2f\n", temperatureC)
// fmt.Printf("%T\n", temperatureC)

// str := fmt.Sprint(80)
// fmt.Println(str)

// fmt.Println("What's your name?")
// var name string
// fmt.Scanln(&name)

// fmt.Printf("Hello, %v\n", name)

// const lowSpeed, highSpeed = 65, 120
// speed := 130
// if speed > highSpeed {
// 	fmt.Println("Speed is too high")
// } else if speed < lowSpeed {
// 	fmt.Println("Speed is too low")
// }

// isWeekend := false
// isHoliday := false

// if isWeekend && isHoliday {
// 	fmt.Println("It's a holiday on a weekend.")
// } else if isWeekend && !isHoliday {
// 	fmt.Println("I's a regular weekend day.")
// } else if !isWeekend && isHoliday {
// 	fmt.Println("It's a weekday  but a holiday.")
// } else {
// 	fmt.Println("It's a regular weekday.")
// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello-world", handleHelloWorld)
	http.HandleFunc("/health", handleHealth)

	addr := "localhost:8000"
	log.Printf("Listening on %s...", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelloWorld(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(w, "Hello World\n")

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	writeResponse(w, "Ok!\n")
}

func writeResponse(w http.ResponseWriter, responseString string) {

	response := []byte(responseString)
	_, err := w.Write(response)
	if err != nil {
		fmt.Println(err)
	}
}
