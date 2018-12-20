package main

import (
	"context"
	"log"

	pb "github.com/aaronflower/shippy-service-user/proto/auth"
	micro "github.com/micro/go-micro"
)

const topic = "user.created"

// define a subscriber to implements Process function.
type subscriber struct {
}

func (sub *subscriber) Process(ctx context.Context, user *pb.User) error {
	log.Println("Picked up a new message")
	log.Println("Sending email to:", user.Name)
	return nil
}

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
	)

	srv.Init()

	micro.RegisterSubscriber(topic, srv.Server(), new(subscriber))

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}
