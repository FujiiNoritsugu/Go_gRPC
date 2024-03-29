package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "Go_gRPC/api/order"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 注文情報の構造体
type Order struct {
	Id     int
	Content  string
}

// DBを使わずに書籍情報を返せるように構造体にデータを保持
var (
	order1 = Order{
		Id:     1,
		Content:  "Content1",
	}
	order2 = Order{
		Id:     2,
		Content:  "Content2",
	}
	orders = []Order{order1, order2}
)

func getOrder(i int32) Order {
	return orders[i-1]
}

// OrderToServerのサービス実装用インターフェースの構造体
type server struct{
	pb. UnimplementedOrderToServer
}

// 自動生成された`order_grpc.pb.go`の'GetOrder`インターフェースを実装。
func (s *server) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.GetOrderResponse, error){
	order := getOrder(in.Id)

	protoOrder := &pb.Order{
		Id: int32(order.Id),
		Content: order.Content,
	}
	return &pb.GetOrderResponse{Order: protoOrder}, nil
}

// 起動するポートを指定
var (
	port = flag.Int("port", 50051, "the port to serve on")
)

func main(){
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPCサーバーを起動
	s := grpc.NewServer()
	// 自動生成された`order_grpc.pb.go`の`RegisterOrderToServer`を呼び出し、サーバーにサービスを登録
	pb.RegisterOrderToServer(s, &server{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}
