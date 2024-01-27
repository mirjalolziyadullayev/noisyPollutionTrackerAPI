package main

import (
	"fmt"
	"net/http"
	"noisyPollutionTrackerAPI/Handlers"
	"time"
)

func main() {
	fmt.Println("server started at: ", time.Now())

	http.HandleFunc("/", Handlers.CityHandler)
	http.ListenAndServe(":8080", nil)
}
