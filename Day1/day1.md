# DAY 1
- Basic syntax
- Go module
- Variable and Declaration

### Basic syntax:
*Learn about the basic syntax of Go, such as how the go programs are executed, package imports, main function, and so on.*

- **package main:** is a special package containing the main function to execute the program
- **import "...":** to import packages for the program
- **func main():**  is a function that does not need a parameter to be passed and the program will execute main function
- **How to execute the program:** use ***go run file.go*** to directly compile the program or use ***go build file.go*** to build a executable file and use ***./exe_file*** to run program

### Go modules
*Go modules are a group of related packages that are versioned and distributed together. They specify the requirements of our project, list all the required dependencies, and help us keep track of the specific versions of installed dependencies.*
- be used to manage the package dependencies in Go
- ***go mod init name_module:*** create a file go.mod contains information about project dependencies and their version management.
- ***go mod tidy:*** to download/update package dependencies
- ***go get package:*** to download package dependencies

### Variable and Declaration
*Variable is the name given to a memory location to store a value of a specific type. Go provides multiple ways to declare and use variables.*

- one way: ***var name type*** => no need to assign values ​​when declaring
- two way: ***var name*** => required to assign value when declaring
- three way: ***name := value*** => required to assign value when declaring can only be used inside functions



