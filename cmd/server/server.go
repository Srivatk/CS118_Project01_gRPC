package main

import (
    "flag"
    "fmt"
    "google.golang.org/grpc"
    whatsup "whatsup/pkg"
)

const (
	DEFAULT_SERVER_PORT        string = "50001"
)


func main() {

    serverPortPtr := flag.String("port", DEFAULT_SERVER_PORT, "chat server port to connect to")
    flag.Parse()

    listen, port, err := whatsup.OpenListener(*serverPortPtr)
    fmt.Printf("Listening on port %s\n", port)

    if err != nil {
        fmt.Println(err)
        return
    }

    whatsupService := whatsup.NewServer()

    realServer := grpc.NewServer(
        grpc.UnaryInterceptor(whatsupService.Interceptor),
    )
    whatsup.RegisterWhatsUpServer(realServer, whatsupService)
    if err := realServer.Serve(listen); err != nil {
        fmt.Printf("failed to serve: %v", err)
    }
}
