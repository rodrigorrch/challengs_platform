package main

import (
	"fmt"
	"log"

	"github.com/rodrigorrch/challengs_platform/application/repositories"
	"github.com/rodrigorrch/challengs_platform/domain"
	"github.com/rodrigorrch/challengs_platform/framework/utils"
)

func main() {

	db := utils.ConnectDB()

	user := domain.User{
		Name:     "Rodrigo",
		Email:    "rodrigo.rrch@gmail.com",
		Password: "123",
	}

	userRepo := repositories.UserRepositoryDb{Db: db}
	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Name, result.Email, result.Token)
}
