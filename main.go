package main

import (
	"context"
	"fmt"
	"log"
	"net"

	dataUser "github.com/knight7024/grpc-go-tutorial/example-data/user"
	pbUser "github.com/knight7024/grpc-go-tutorial/proto/user"
	"google.golang.org/grpc"
)

const networkName string = "tcp"
const portNumber string = "9001"

type userServer struct {
	pbUser.UserServer
}

func (s *userServer) GetAddressBook(ctx context.Context, req *pbUser.GetAddressBookRequest) (*pbUser.AddressBook, error) {
	addressbookMessages := make([]*pbUser.Person, len(dataUser.People))
	for i, u := range dataUser.People {
		addressbookMessages[i] = u
	}

	return &pbUser.AddressBook{
		People: addressbookMessages,
	}, nil
}

func main() {
	// tcp 9001 포트로 요청받는 Listener 생성
	lis, err := net.Listen(networkName, fmt.Sprintf(":%s", portNumber))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// gRPC 서버를 생성해서
	grpcServer := grpc.NewServer()
	pbUser.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC server on %s port", portNumber)
	// 해당 Listener로 시작
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
