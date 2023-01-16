package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/kikytokamuro/grpc_py_go_example/dir_watcher"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5030, "The server port")
)

type server struct {
	pb.UnimplementedDirWatcherServer
}

func (s *server) Do(ctx context.Context, in *pb.DirWatchRequest) (*pb.DirWatchResponse, error) {
	files, err := ioutil.ReadDir(in.GetDirectory())
	if err != nil {
		return &pb.DirWatchResponse{}, err
	}

	content := make([]string, len(files))

	for _, file := range files {
		if file.IsDir() {
			content = append(content, fmt.Sprintf("%s/", file.Name()))
		} else {
			content = append(content, file.Name())
		}
	}

	return &pb.DirWatchResponse{Content: content}, nil
}

func main() {
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterDirWatcherServer(srv, &server{})

	log.Printf("Server listening at %v", listen.Addr())

	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
