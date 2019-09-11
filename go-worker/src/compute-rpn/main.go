package main

import (
	"compute-rpn/rpn"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ComputeTask Simple struct to handle json request
type ComputeTask struct {
	Exp string `json:"exp"`
}

// ComputeResult Its better to make json object as strings
type ComputeResult struct {
	Result string `json:"result"`
	Time   string `json:"time"`
	Status string `json:"status"`
}

// in gorilla mux every request is a separate goroutine. So single expression - single gorutine
func computeRPN(w http.ResponseWriter, req *http.Request) {
	computeRPN := rpn.RPN{}
	var newTask ComputeTask
	//struct for jsonResult
	jsonStruct := &ComputeResult{Result: "", Time: "", Status: ""}
	//flag for writing good response
	decodeErr := json.NewDecoder(req.Body).Decode(&newTask)
	// calculate result
	startTime := time.Now().UnixNano()
	computedResult, err := computeRPN.CalcRPN(newTask.Exp)
	elapsedTime := time.Now().UnixNano() - startTime
	if decodeErr != nil {
		jsonStruct.Status = "decode_error"
		jsonStruct.Result = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else if err != nil {
		jsonStruct.Status = "compute_error"
		jsonStruct.Result = err.Error()
		w.WriteHeader(http.StatusBadRequest)
	} else {
		jsonStruct.Status = "ok"
		jsonStruct.Result = strconv.Itoa(computedResult)
		jsonStruct.Time = strconv.FormatInt(elapsedTime, 10) + "ns"
		w.WriteHeader(http.StatusOK)
	}
	jsonResult, marshalErr := json.Marshal(jsonStruct)
	fmt.Println(marshalErr)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResult))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/compute", computeRPN).Methods("POST")
	fmt.Println("Compute Server started at 8080")
	http.ListenAndServe(":8080", r)
}
