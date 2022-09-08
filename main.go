package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

type InputDataMean struct {
	Data []int `json:"data"`
}

type OutputDataMean struct {
	Data   []int   `json:"data"`
	Result float64 `json:"result"`
}
type OutputDataMedian struct {
	Data   []int `json:"data"`
	Result int   `json:"result"`
}

func meanhandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" { // <1>
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // <2>
		return                                                           // <3>
	}
	if r.Header.Get("Content-Type") != "application/json" { // <4>
		http.Error(w, "Content-Type not allowed", http.StatusBadRequest) // <5>
		return
	}
	inputData := &InputDataMean{}                    // <6>
	err := json.NewDecoder(r.Body).Decode(inputData) // <7>
	if err != nil {                                  // <8>
		http.Error(w, "json.NewDecoder", http.StatusBadRequest) // <9>
		return
	}

	sum := 0                           // <10>
	for _, v := range inputData.Data { // <11>
		sum += v // <12>
	}

	outputDataMean := &OutputDataMean{ // <13>
		Data:   inputData.Data,                              // <14>
		Result: float64(sum) / float64(len(inputData.Data)), // <15>
	}

	w.Header().Set("Content-Type", "application/json") // <16>
	json.NewEncoder(w).Encode(outputDataMean)          // <17>
}

func medianhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" { // <1>
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // <2>
		return
	}
	if r.Header.Get("Content-Type") != "text/csv" { // <3>
		http.Error(w, "Content-Type not allowed", http.StatusBadRequest) // <4>
		return
	}
	csvData := csv.NewReader(r.Body) // <5>
	csvData.TrimLeadingSpace = true  // <6>
	rec, err := csvData.Read()       // <7>
	if err != nil {                  // <8>
		http.Error(w, "Parse error", http.StatusBadRequest) // <9>
		return
	}
	var data []int       // <10>
	for i := range rec { // <11>
		ival, err := strconv.Atoi(rec[i]) // <12>
		if err != nil {                   // <13>
			http.Error(w, "strconv.Atoi error", http.StatusBadRequest) // <14>
			return
		}
		data = append(data, ival) // <15>
	}
	sort.Ints(data)   // <16>
	median := 0       // <17>
	size := len(data) // <18>
	if size%2 == 0 {  // <19>
		median = (data[size/2] + data[size/2-1]) / 2 //  <20>
	} else {
		median = data[size/2] // <21>
	}

	outputDataMedian := &OutputDataMedian{ // <22>
		Data:   data,   // <23>
		Result: median, // <24>
	}

	w.Header().Set("Content-Type", "application/json") // <25>
	json.NewEncoder(w).Encode(outputDataMedian)        // <26>
}
func main() { //

	log.Print("starting server...")
	http.HandleFunc("/", echohandler)              // <1>
	http.HandleFunc("/stat/mean", meanhandler)     // <2>
	http.HandleFunc("/stat/median", medianhandler) // <3>

	// Determine port for HTTP service.
	port := os.Getenv("PORT") // <4>
	if port == "" {           // <5>
		port = "8085"                             // <6>
		log.Printf("defaulting to port %s", port) // <7>
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)                   // <8>
	if err := http.ListenAndServe(":"+port, nil); err != nil { // <9>
		log.Fatal(err) // <10>
	}
}
func echohandler(w http.ResponseWriter, r *http.Request) { //
	name := os.Getenv("NAME") // <1>
	if name == "" {           // <2>
		name = "World" // <3>
	}
	fmt.Fprintf(w, "Hello %s!\n", name) // <4>
}
