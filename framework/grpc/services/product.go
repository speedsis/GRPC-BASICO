package services

import (
	"context"
	"github.com/speedsis/code-grpc/domain"
	"github.com/speedsis/code-grpc/framework/grpc/pb"
	"time"
)

type ProductGrpcServer struct {

	pb.UnimplementedProductServiceServer

	Products *domain.Products
}

func (p *ProductGrpcServer) CreateProduct(ctx context.Context, in *pb.Product) (*pb.ProductResult, error) {

	product := domain.NewProduct()
	product.Name = in.Name

	p.Products.Add(product)

	return &pb.ProductResult{
		Id: product.ID,
		Name: product.Name,
	}, nil
}

func (p *ProductGrpcServer) List(req *pb.Empty, stream pb.ProductService_ListServer)  error {

	for _, product := range p.Products.Product {
		time.Sleep(time.Second * 2)
		stream.Send(&pb.ProductResult{
			Id: product.ID,
			Name: product.Name,
		})
	}
	return nil
}

func NewProductGrpcServer(products *domain.Products) *ProductGrpcServer {


	return &ProductGrpcServer{
		Products: products,
	}
}
