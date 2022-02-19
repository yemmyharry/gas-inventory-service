package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	adapter "github.com/yemmyharry/gas-inventory-service/internal/adapter/api/resource"
	repository "github.com/yemmyharry/gas-inventory-service/internal/adapter/repositories/mongodb/resource"
	"github.com/yemmyharry/gas-inventory-service/internal/core/helper"
	"time"
)

func main() {
	helper.InitializeLog()
	address, port, _, _, _, _ := helper.LoadConfig()
	router := gin.New()
	db := &repository.MongoRepository{}
	db.Init()
	s := &adapter.Server{
		Inventory: db,
		Route:     router,
	}
	s.Routes(router)

	fmt.Println("Service running on " + address + ":" + port)
	helper.LogEvent("Info", fmt.Sprintf("Started PlatformServiceApplication on "+address+":"+port+" in "+time.Since(time.Now()).String()))
	_ = s.Route.Run(":" + port)
}
