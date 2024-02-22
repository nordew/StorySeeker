package main

import (
	"fmt"
	"github.com/nordew/StorySeeker/pkg/seeker"
	"log"
	"net/http"
	"os"
)

const (
	COOKIE    = ``
	userAgent = "USER_AGENT = \"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Instagram 105.0.0.11.118 (iPhone11,8; iOS 12_3_1; en_US; en-US; scale=2.00; 828x1792; 165586599)"
)

func main() {
	client := http.Client{}

	s := seeker.NewSeeker(client, userAgent)

	if len(os.Args) < 2 {
		log.Fatal(" go run main.go <username>")
	}
	username := os.Args[1]

	stories, err := s.Get(COOKIE, username)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	for _, story := range stories {
		fmt.Printf("Story: %s", story)
	}
}
