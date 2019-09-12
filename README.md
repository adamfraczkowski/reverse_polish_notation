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
	"exp":"12 2 3 4 * 10 5 / + * +"
}
```

HTTP Server run on port 8080. Based on gorilla/mux package. Every request is concurrent, so there is no need to additionaly parallelise computations

## python-api

Python-api is responsible for logging and making task queue. Application runs simple Flask service on 8081 port. Had single API endpoint `/queue`.

Example JSON request:

```
["12 2 3 4 * 10 5 / + * +","7 2 3 4 * 10 5 / + * +","4 2 3 4 * 10 5 / + * +"]
```
And response:
```
[
  {
    "result": "40",
    "time": "0ns"
  },
  {
    "result": "35",
    "time": "0ns"
  },
  {
    "result": "32",
    "time": "0ns"
  }
]
```

Generate logs both on stdout and to `python.log` file. It uses python logging mechanism

## ruby-endpoint

TODO

# RUNNING PROJECT

## Requirements

* Python 3.6
* Ruby 2.x
* pip for python3
* Bundler Gem for Ruby

## Instalation

* for python go to `python-api` folder and run command:
`pip install -r requirements.txt`
* for ruby go t `ruby-endpoint` and run command:
`bundle install`

## Execute

You can manually run every project in order:
* go-worker
* python-api
* ruby-endpoint

Depends on OS and your env you need execute three commands in separate terminals:

```
go run go-worker/src/compute-rpn/main.go 
python3 python-api/server.py 
ruby ruby-endpoint/main.rb
``` 

There's also `insomnia` workspace to test server application.