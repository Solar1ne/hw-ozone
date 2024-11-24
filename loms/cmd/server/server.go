package main

import (
	"awesomeProject4/loms/internal/app/server"
	"awesomeProject4/loms/pkg/api/loms/v1"
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	grpcAddr := ":50051"
	httpAddr := ":8080"

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	lomsServer := server.NewServer()
	loms.RegisterLOMSServiceServer(grpcServer, lomsServer)

	go func() {
		log.Printf("starting gRPC server on %s", grpcAddr)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	gwServer := &http.Server{ //ххххххххххххххххххххххххххххх
		Addr: httpAddr,
	}

	go func() {
		log.Printf("starting HTTP server on %s", httpAddr)
		if err := gwServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1) //
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("shutting down servers")
	gwServer.Shutdown(context.Background()) //  зачем хендлить хххх
	grpcServer.GracefulStop()
}
