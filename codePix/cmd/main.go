package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/jonascesconetto/imersao-fullcycle/codePix/application/grpc"
	"github.com/jonascesconetto/imersao-fullcycle/codePix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
