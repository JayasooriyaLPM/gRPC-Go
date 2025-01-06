package main

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/text/message"
	"google.golang.org/grpc"

	"jayasooriyalpm/go-grpc/chat"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}
	defer conn.Close()

	c:= chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello From the Client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from the server: %s", response.Body)
}
