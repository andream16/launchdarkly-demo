package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	ld "gopkg.in/launchdarkly/go-client.v4"
)

func stringToPtr(s string) *string {
	return &s
}

func main() {

	client, err := ld.MakeClient("sdk-c1be7e8b-a422-4f06-a97f-965a9f6b0381", 5*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	for i := 0; i <= 10; i++ {

		id := uuid.New().String()
		user := ld.User{
			Key: &id,
		}

		showFeature, err := client.BoolVariation("amex", user, false)
		if err != nil {
			log.Fatal(err)
		}

		if showFeature {
			fmt.Println(fmt.Sprintf("Showing your feature to user: %s", id))
		} else {
			fmt.Println(fmt.Sprintf("Not showing your feature to user: %s", id))
		}
	}

}
