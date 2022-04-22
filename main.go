package main

import (
	"context"
	"fmt"
	"log"
	"net"

	livescore "grpc-demo/proto/go/proto/livescore"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var matches livescore.ListMatchesResponse

type liveScoreServer struct {
	livescore.UnimplementedScoreServiceServer
}

func (lss *liveScoreServer) ListMatches(ctx context.Context, req *livescore.ListMatchesRequest) (*livescore.ListMatchesResponse, error) {

	match := &livescore.MatchScoreResponse{
		Score: "4:1",
		Live:  true,
	}

	matches.Scores = append(matches.Scores, match)
	return &matches, nil
}

const addr = "localhost:50004"

func main() {
	//create tcp connection on port 50004
	conn, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("tcp connection err: ", err.Error())
	}
	//create grpc server
	grpcServer := grpc.NewServer()

	server := liveScoreServer{}
	//register our livescore service
	livescore.RegisterScoreServiceServer(grpcServer, &server)

	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server at : ", addr)
	//serve our connection
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
