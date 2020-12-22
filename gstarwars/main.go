package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	grpc "google.golang.org/grpc"
)

func main() {
	// Build the server
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	RegisterStarwarsServer(server, &starwars{})

	go func() {
		fmt.Println("Starting Go server on port 9001")
		err = server.Serve(listener)
		if err != nil {
			panic(err)
		}
	}()

	// Call the C# server on port 9002
	conn, err := grpc.Dial("localhost:9002", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := NewStarwarsClient(conn)

	fmt.Println("Enter the command to call the C# server...")
	fmt.Println("	c/name- search characters")
	fmt.Println("	h/name - get hero by episode")
	fmt.Println("	k/name - create review")
	fmt.Println("	r/name - get reviews by episode")
	fmt.Println("	x - exit")
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\r\n")
		fmt.Println(text)
		s := strings.Split(text, "/")
		cmd, name := s[0], s[1]
		if cmd == "x" {
			break
		}

		switch cmd {
		case "c":
			req := SearchCharacterRequest{Name: name}
			resp, err := client.SearchCharacter(context.Background(), &req)
			if err != nil {
				panic(err)
			}

			fmt.Println(resp)
			break
		case "h":
			//fmt.Println(name)
			value, ok := Episode_value[name]
			//fmt.Println(value)
			if !ok {
				fmt.Println("invalid enum value")
				break
			}

			req := GetHeroRequest{Episode: Episode(value)}
			resp, err := client.GetHero(context.Background(), &req)
			if err != nil {
				panic(err)
			}

			fmt.Println(resp)
			break
		case "k":
			if len(s) != 4 {
				fmt.Println("invalid input")
				break
			}

			stars, err := strconv.Atoi(s[2])
			if err != nil {
				// handle error
				fmt.Println(err)
				break
			}

			value, ok := Episode_value[name]
			if !ok {
				fmt.Println("invalid enum value")
				break
			}

			episode := Episode(value)
			req := Review{Stars: int32(stars), Commentary: s[3], Episode: episode}
			resp, err := client.AddReview(context.Background(), &req)
			if err != nil {
				panic(err)
			}

			fmt.Println(resp)
			break
		case "r":
			value, ok := Episode_value[name]
			if !ok {
				fmt.Println("invalid enum value")
				break
			}

			req := GetReviewsRequest{Episode: Episode(value)}
			resp, err := client.GetReviews(context.Background(), &req)
			if err != nil {
				panic(err)
			}

			fmt.Println(resp)
			break
		default:
			break
		}
	}

	// Wait to exit
	fmt.Scanln()
}
