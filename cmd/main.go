package main

import (
	"github.com/mayerkv/go-candidates/domain"
	"github.com/mayerkv/go-candidates/grpc-service"
	"github.com/mayerkv/go-candidates/repository"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	candidateRepository := repository.NewInMemoryCandidateRepository()
	candidateService := domain.NewCandidateService(candidateRepository)
	srv := grpc_service.NewCandidateServiceImpl(candidateService)

	if err := runGrpcServer(srv); err != nil {
		log.Fatal(err)
	}
}

func runGrpcServer(server grpc_service.CandidatesServiceServer) error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpc_service.RegisterCandidatesServiceServer(grpcServer, server)

	return grpcServer.Serve(lis)
}
