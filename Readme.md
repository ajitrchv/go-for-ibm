# Notes on Go-for-IBM

## How do we run the code inside the project?

```
1. go mod init example.com/m
2. go mod tidy
3. go run main.go

"hello world!"
```


```go Build```: compiles the program.

```go run``` : compiles and then immediately executes the program.

```go fmt``` : automatically formats all codes inside thye current dir.

```go install``` : compiles and installs a package

```go get```: downloads a raw source code of someone else's package

```go test``` : tests are run associated with the current project``

## What does the package main means?

Two types of packages :
    - Executable: Generates a file that we can run.
    - Reusable: Code used as helpers. Place to put reuasable logic.

If the package is ```main```, then thaat is an executable.