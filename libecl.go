package main

import (
	"fmt"
	"net/http"
)

import "C"

// To build: go build -buildmode=c-archive
// Or to debug, go build -buildmode=c-archive -x -work
// See https://golang.org/cmd/go/

//export ActivateFeature
func ActivateFeature(featureId string) (retVal int) {

	// Simulate making an https call - in this case just getting an https
	// example home page
	tr := &http.Transport{
		// Can define cert pool here as an option if we need to ship certs...
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://example.com/")
	if err != nil {
		fmt.Println("failed to make HTTPS request")
		return -1
	}
	defer resp.Body.Close()

	fmt.Printf("activated feature %s in Go!\n", featureId)
	fmt.Printf("Response from server was: %s \n", resp.Status)
	retVal = 3
	return
}

func main() {
	// This is ignored when making a library but it still needs to be here for
	// some reason
}
