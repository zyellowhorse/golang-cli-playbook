## Purpose of this module
This module will provide information about retrieving the environment information in `go`.
Then you can examining the `go` environment manually or through a tool like `grep`.

## Verify locally
To test this module locally:
* Open up a terminal at the project root
* Run command `cd module1` to change to the `module1` directory
* Run command `go test -run Module1` to run all tests for module 1, or
* Run command `go test -v -run Module1` to run all tests for module 1 with verbose information


## Task 1: Print the environment information
First, refer to the [Print Go environment information](https://golang.org/cmd/go/#hdr-Print_Go_environment_information) section in the official documentation.

Then, type in the command for printing the environment information.
In this task, we only need to specify the sub-command after `go`, and do not need to specify any flags after it.

## Task 2: Print the environment information and in JSON format
By default, `go env` outputs as a shell script on Linux / MacOS and a batch script on Windows.

[JavaScript Object Notation (JSON)](https://en.wikipedia.org/wiki/JSON) is a common format for data exchanging.

Recall what you have read in the above documentation and check flags that are available to the `go env` command.
What is the flag to specify the output to be in JSON format?

Write the complete command (go + sub-command + flag) under your answer for Task 1.

## Extra help
Here are the commands to get more information on `go env` and `Go` environment:
- `go help env`
- `go help environment`
