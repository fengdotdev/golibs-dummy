package main

import (
	"fmt"

	"github.com/fengdotdev/golibs-helperfuncs/ident"
	"github.com/fengdotdev/golibs-helperfuncs/web"
)

func main() {
	fmt.Println("Hello, world!")
	identity := ident.DeterministicUUID("foo", "bar")

	fmt.Println("Identity:", identity)
	

	

	_, err := web.GetRemoteResource("https://pkg.go.dev/github.com/fengdotdev/golibs-helperfuncs@v1.0.0?tab=versions")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}
}
