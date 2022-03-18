package main

import (
	"fmt"
	adapter "gas-inventory-service/internal/adapter/api/resource"
	repository "gas-inventory-service/internal/adapter/repositories/mongodb/resource"
	"gas-inventory-service/internal/core/helper"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	helper.InitializeLog()
	address, port, _, dbHost, dbName, _ := helper.LoadConfig()
	router := gin.New()
	db := &repository.MongoRepository{}
	db.Init(dbHost, dbName)
	s := &adapter.Server{
		Inventory: db,
		Route:     router,
	}
	s.Routes(router)

	fmt.Println("Service running on " + address + ":" + port)
	helper.LogEvent("Info", fmt.Sprintf("Started PlatformServiceApplication on "+address+":"+port+" in "+time.Since(time.Now()).String()))
	_ = s.Route.Run(":" + port)
}
