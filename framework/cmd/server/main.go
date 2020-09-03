package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rodrigorrch/challengs_platform/application/repositories"
	"github.com/rodrigorrch/challengs_platform/application/usecases"
	"github.com/rodrigorrch/challengs_platform/framework/pb"
	"github.com/rodrigorrch/challengs_platform/framework/servers"
	"github.com/rodrigorrch/challengs_platform/framework/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {

	db = utils.ConnectDB(os.Getenv("env"))

	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	userServer := setUpUserServer()
	challengeServer := setUpChallengeServer()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userServer)
	pb.RegisterChallengeServiceServer(grpcServer, challengeServer)
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}

func setUpUserServer() *servers.UserServer {
	userServer := servers.NewUserServer()
	userRepository := repositories.UserRepositoryDb{Db: db}
	userServer.UserUseCase = usecases.UserUseCase{UserRepository: userRepository}
	return userServer
}

func setUpChallengeServer() *servers.ChallengeServer {
	challengeServer := servers.NewChallengeServer()
	challengeRepository := repositories.ChallengeRepositoryDb{Db: db}
	challengeServer.ChallengeUseCase = usecases.ChallengeUseCase{ChallengeRepository: challengeRepository}
	return challengeServer
}
