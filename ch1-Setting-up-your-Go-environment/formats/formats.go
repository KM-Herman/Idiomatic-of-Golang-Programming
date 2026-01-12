package main

import "fmt"

func main() {

	//go fmt
	fmt.Println("Hello, World!")

	// go vet this detects errors of syntactically valid codes but likely incorrect
	fmt.Printf("Hello, %s!\n") // beacause we are not passing any arguments to the printf function we will get an error
	//.\formats.go:11:21: fmt.Printf format %s reads arg #1, but call has 0 args

}
