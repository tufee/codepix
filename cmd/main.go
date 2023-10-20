package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/tufee/codepix/application/grpc"
	"github.com/tufee/codepix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
