#!/usr/bin/ruby
require 'json'
require 'httparty'

QUEUE_SERVER = "http://localhost:8081/queue"

# example info message
expression_array = []
puts "Put number of expressions and expressions, separated by \\n:"
input_num = $stdin.gets.chomp
# Check number of inputs
exp_num = input_num.to_i
# Get expressions
(1..exp_num).each { |i|
    temporary_input = $stdin.gets.chomp
    expression_array << temporary_input
}
#prepare json struct
jsonStruct = JSON[expression_array]
#send data to queue server
result_of_queue_request = HTTParty.post(QUEUE_SERVER,:body => jsonStruct,:headers => {'Content-Type' => 'application/json'})
#parsing sorted result to output
(0..result_of_queue_request.length-1).each { |i|
    puts result_of_queue_request[i]["result"] + ","+ result_of_queue_request[i]["time"] 
}
