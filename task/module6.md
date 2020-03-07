## Purpose of this module
This module will teach you how to use `go doc` to prints the documentation for a package, a function, a variable, or a method.
We will learn the `go doc` command and the possible parameters available to it


## Verify locally
To test this module locally:
* Open up a terminal at the project root
* Run command `cd module6` to change to the `module6` directory
* Run command `go test -run Module6` to run all tests for module 6, or
* Run command `go test -v -run Module6` to run all tests for module 6 with verbose information

## Task 1
A `go` source file named `module6.go` is provided, 
but the documentation for a function (`FunctionForModule6GoDoc`)is missing.

Complete the task by adding a comment in the source file.


## Task 2
To see the documentation added in Task 1, we need to specify a parameter to the `go doc` command.

Read the `go doc` command document, and provide a parameter in the following format:
```
go doc sym.methodOrField
```

Does the command show the documentation added in Task 1?


## Extra help
Here is the command to get more information on `go doc`: 
- `go help doc`

