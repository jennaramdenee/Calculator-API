# Calculator API

### Purpose

A simple HTTP API that uses Golang to perform basic calculations, as input by the user (currently supports add, multiply operators).

### How to Use

1. Fork this repository
2. Clone to your local machine
3. Run `go get` to ensure all required modules
4. Run `cd src` to navigate to the `src` folder
5. Run `go build` to build the program
6. Run `./src` to start the program
7. Navigate to `http://localhost:8080/` to use the API
  - For addition, visit `http://localhost:8080/add/x,y,z` where `x`, `y` and `z` are numbers to add
  - For multiplcation, visit `http://localhost:8080/multiply/x,y,z` where `x`, `y` and `z` are numbers to multiply
