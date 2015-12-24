package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func PrintUsage() {
	fmt.Println("go run generator/main.go count")
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		PrintUsage()
		os.Exit(1)
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, "count must be integer")
		os.Exit(1)
	}
	if count <= 0 {
		fmt.Fprintln(os.Stderr, "count must be greater than 0")
		os.Exit(1)
	}

	declOfReturnValuesList := make([]string, count)
	returnValuesList := make([]string, count)
	underscores := make([]string, count)

	for i := 0; i < count; i++ {
		declOfReturnValuesList[i] = fmt.Sprintf("arg%d", i)
		returnValuesList[i] = "0"
		underscores[i] = "_"
	}

	t := template.Must(template.ParseFiles("template.go"))
	s := struct {
		DeclOfReturnValuesList, ReturnValuesList, Underscores string
	}{
		strings.Join(declOfReturnValuesList, ", "),
		strings.Join(returnValuesList, ", "),
		strings.Join(underscores, ", "),
	}
	t.Execute(os.Stdout, s)
}
