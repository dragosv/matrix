# Matrix

[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#wip) [![Build Status](https://travis-ci.com/dragosv/matrix.svg?branch=master)](https://travis-ci.com/dragosv/matrix)
[![codecov](https://codecov.io/gh/dragosv/matrix/branch/master/graph/badge.svg)](https://codecov.io/gh/dragosv/matrix)
[![Go Report Card](https://goreportcard.com/badge/github.com/dragosv/matrix)](https://goreportcard.com/report/github.com/dragosv/matrix)

Basic web server written in GoLang that performs operations on a matrix

## Build and Run

To build and run the web server the following commands needs to executed at the terminal 

```
go mod download
go mod verify
go build .
go run .
``` 

## Api
The description of the API can be accessed while the server is running at http://localhost:8080/swagger/index.html

### Echo 
Return the matrix as a string in matrix format.

```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// Expected output
1,2,3
4,5,6
7,8,9
``` 
### Invert
Return the matrix as a string in matrix format where the columns and rows are inverted
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/invert"

// Expected output
1,4,7
2,5,8
3,6,9
``` 
### Flatten
Return the matrix as a 1 line string, with values separated by commas.
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/flatten"

// Expected output
1,2,3,4,5,6,7,8,9
``` 
### Sum
Return the sum of the integers in the matrix. Will return bad request if the addition overflows.
    
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/sum"

// Expected output
45
``` 
### Multiply
Return the product of the integers in the matrix. Will return bad request if the multiplication overflows. 
    
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/multiply"

// Expected output
362880
``` 

## Organization

The http handlers are located in the handlers folder while the corresponding business commands are in the commands folder