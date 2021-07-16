package main

import (
	"CarLibrary/score/config"
	"CarLibrary/score/model"
	"CarLibrary/score/score"
	pb "github.com/CarLibrary/proto/score"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {

	config.Config()
	model.InitMYSQL()

	lis, err := net.Listen("tcp", os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterScoreServiceServer(s, &score.ScoreServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
