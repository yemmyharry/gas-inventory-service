package main

import (
	"fmt"
	adapter "gas-inventory-service/internal/adapter/api/resource"
	repository "gas-inventory-service/internal/adapter/repositories/mongodb/resource"
	"gas-inventory-service/internal/core/helper"
	services "gas-inventory-service/internal/core/services/resource"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	helper.InitializeLog()
	address, port, _, dbHost, dbName, _ := helper.LoadConfig()
	router := gin.New()
	db := &repository.MongoRepository{}
	db.Init(dbHost, dbName)
	service := services.NewService(db)
	handler := adapter.NewHTTPHandler(service, router)
	handler.Routes(router)

	fmt.Println("Service running on " + address + ":" + port)
	helper.LogEvent("Info", fmt.Sprintf("Started PlatformServiceApplication on "+address+":"+port+" in "+time.Since(time.Now()).String()))
	_ = handler.Route.Run(":" + port)
}
