package server

import (
	"github.com/speedsis/code-grpc/domain"
	"github.com/speedsis/code-grpc/framework/grpc/pb"
	"github.com/speedsis/code-grpc/framework/grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var ProductList = domain.NewProducts()

func StartGrpcServer() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cold no connected ", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	productService := services.NewProductGrpcServer(ProductList)
	pb.RegisterProductServiceServer (grpcServer, productService)

	log.Println("Grpc server has been started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("", err)
	}

}
