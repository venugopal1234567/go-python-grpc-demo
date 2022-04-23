package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	livescore "grpc-demo/proto/go/proto/livescore"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var msgChan chan string
var matches livescore.ListMatchesResponse

type liveScoreServer struct {
	livescore.UnimplementedScoreServiceServer
}

func (lss *liveScoreServer) ListMatches(ctx context.Context, req *livescore.ListMatchesRequest) (*livescore.ListMatchesResponse, error) {

	match := &livescore.MatchScoreResponse{
		Score: "4:1",
		Live:  true,
	}

	if msgChan != nil {
		msg := "Hello this is an event"
		msgChan <- msg
	}

	matches.Scores = append(matches.Scores, match)
	return &matches, nil
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Client connected")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	msgChan = make(chan string)

	defer func() {
		close(msgChan)
		msgChan = nil
		fmt.Println("Client closed connection")
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("Could not init http.Flusher")
	}

	for {
		select {
		case message := <-msgChan:
			fmt.Println("case message... sending message")
			fmt.Println(message)
			fmt.Fprintf(w, "data: %s\n\n", message)
			flusher.Flush()
		case <-r.Context().Done():
			fmt.Println("Client closed connection")
			return
		}
	}

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

	//Create http server
	router := http.NewServeMux()

	router.HandleFunc("/event", sseHandler)
	go func() {
		fmt.Println("Starting http server at : localhost:3500")
		if err := http.ListenAndServe(":3500", router); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		fmt.Println("Starting gRPC server at : ", addr)
		//serve our connection
		if err := grpcServer.Serve(conn); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown", s)
		grpcServer.GracefulStop()
		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	wg.Wait()
	log.Println("clean shutdown")

}
