package main

import (
	"log"
	"time"

	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	accountProto "github.com/timkellogg/store/account/protos/account"
)

var accountRepository = AccountRepository{}
var config = Config{}

func init() {
	config.Read()

	accountRepository.Server = config.Server
	accountRepository.Database = config.Database
	accountRepository.Connect()
}

func main() {
	service := k8s.NewService(
		micro.Name("go.micro.srv.accounts"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*30),
	)

	accountProto.RegisterAccountsHandler(service.Server(), new(Server))

	service.Init()

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
