package main

import (
	"fmt"
	"log"
	"time"

	ld "gopkg.in/launchdarkly/go-client.v4"
)

func main() {

	client, err := ld.MakeClient("sdk-c1be7e8b-a422-4f06-a97f-965a9f6b0381", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	key := "andream16"

	user := ld.User{
		Key: &key,
	}

	showFeature, err := client.BoolVariation("amex", user, false)
	if err != nil {
		log.Fatal(err)
	}

	if showFeature {
		fmt.Println("Showing your feature")
	} else {
		fmt.Println("Not showing your feature")
	}

}
