package main

import "fmt"

// Run is going to be responsible for
// instantiation and startup our go application
func Run() error {
	fmt.Println("starting our application")
	return nil
}

func main() {
	fmt.Println("server")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
