package main

func f() ({{.DeclOfReturnValuesList}} int) {
	return {{.ReturnValuesList}}
}

func main() {
  {{.Underscores}} = f()
}
