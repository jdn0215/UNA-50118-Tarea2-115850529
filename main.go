package main

import (
    pb "github.com/jdn0215/UNA-50118-Tarea2-115850529/booksapp"
    "google.golang.org/grpc"
    "log"
    "net"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    lis, err := net.Listen("tcp", ":"+port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterBookInfoServer(s, &server{})

    log.Printf("Starting gRPC listener on port " + port)

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}