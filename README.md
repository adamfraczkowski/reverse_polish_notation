# reverse_polish_notation

Simple demo project for evaluate Reverse Polish Notation. For real there are three application:

* go-worker - make calculation of provided task. (Golang)
* python-api - simple API for getting requests and log output to file (Python)
* ruby-endpoint - simple script for parsing user input and outputs result

# Architecture 

## go-worker

It's simple computation worker. Calculate RPN expressions. In result, it return status of computation, exact result and calculation time in JSON format. Example of correct result response:
```
{
  "result": "40",
  "time": "3.1109995s",
  "status": "ok"
}
```

Example of bad request response:
```
{
  "result": "Incorrect phrase in expression d12",
  "time": "",
  "status": "compute_error"
}
```
If status 400 received, errors occurs. There are two types of errors:
* compute_error - happens during compute expression. Check result for more verbose messages
* decode_error -happens during request json encoding. Check result for more verbose messages

Example of request, to calculate expression:
```
{
	"exp":"d12 2 3 4 * 10 5 / + * +"
}
```

HTTP Server run on port 8080. Based on gorilla/mux package. Every request is concurrent, so there is no need to additionaly parallelise computations

## python-api

TODO

## ruby-endpoint

TODO